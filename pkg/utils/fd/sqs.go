package fd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/samber/lo"
	"io"
	"os"
	"strings"
	"syscall"
)

//type Processor interface {
//	Setup(input Input, output Output)
//	Run() error
//}
//
//type Process struct {
//	Processor Processor
//}
//
//func setupFds(cfg aws.Config, name string) (io.ReadCloser, io.WriteCloser, func()) {
//	var stdin io.ReadCloser
//	var pid int
//	if utils.IsTTY(os.Stdin) {
//		L.Debug.Println("stdin: terminal")
//		stdin = os.Stdin
//	} else {
//		conf := ReadConf(os.Stdin)
//		stdinArn = conf.Url
//		pid = *conf.Pid
//
//		L.Debug.Println("stdin:", conf.Url)
//		stdin = lo.Must(utils.NewPipe(context.TODO(), conf.Url))
//	}
//
//	stdoutArn, stdoutUrl := SetupSqsFd(cfg, name, "stdout")
//
//	stdout := lo.Must(utils.NewPipe(context.TODO(), stdoutUrl))
//
//	var read func()
//	if utils.IsTTY(os.Stdout) {
//		L.Debug.Println("stdout: terminal")
//		stdout = os.Stdout
//
//		L.Debug.Printf("stdout: %s -> terminal\n", stdoutArn)
//
//		read = func() {
//			lo.Must(io.Copy(os.Stdout, stdout))
//			err := syscall.Kill(pid, 9)
//			if err != nil {
//				L.Error.Println("sending sigkill to:", pid)
//			}
//			L.Debug.Println("sent sigkill to:", pid)
//		}
//	} else {
//		arn, url := SetupSqsFd(cfg, name, "stdout")
//		L.Debug.Println("stdout:", arn)
//		stdout = lo.Must(utils.NewPipe(context.TODO(), url))
//
//		WriteConf(os.Stdout, WrapperConfig{
//			Url: url,
//			Arn: arn,
//			Pid:       aws.Int(os.Getpid()),
//		})
//	}
//
//	return stdin, stdout, read
//}

func SetupSqsFd(cfg aws.Config, name, fd string) (arn string, uri string) {
	L.Debug.Println("creating sqs queue:", name, fd)
	client := sqs.NewFromConfig(cfg)

	queueName := fmt.Sprintf("%s-%s", name, fd)
	resp := lo.Must(client.CreateQueue(context.TODO(), &sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
	}))
	L.Debug.Println("created sqs queue:", *resp.QueueUrl)

	// https://sqs.us-east-1.amazonaws.com/123456789012/cat-msh-0a0c3f9d7adb3006a923c4fd8884951b068f44b5-stdout3
	p := strings.Split(*resp.QueueUrl, "/")
	domain := p[2]
	account := p[3]
	region := strings.Split(domain, ".")[1]

	return fmt.Sprintf("arn:aws:sqs:%s:%s:%s", region, account, queueName), *resp.QueueUrl
}

func NewInputPipe(cfg *Config, remote bool) *InputPipe {
	pipe := &InputPipe{}
	if cfg.WrapperConfig.StdoutLocal {
		L.Debug.Println("stdin: terminal")
		pipe.ReadCloser = os.Stdin
	} else {
		L.Debug.Println("stdin:", cfg.StdoutUrl)
		pipe.Url = cfg.WrapperConfig.StdoutUrl
		pipe.Arn = cfg.WrapperConfig.StdoutArn
		pipe.ReadCloser = lo.Must(NewPipe(context.TODO(), cfg.WrapperConfig.StdoutUrl))
	}

	return pipe
}

type InputPipe struct {
	io.ReadCloser
	Arn string
	Url string
	pid Config
}

func NewRemoteOutputPipe(cfg *Config) *OutputPipe {
	pipe := &OutputPipe{}

	pipe.Arn, pipe.Url = SetupSqsFd(cfg.AwsCfg, cfg.Name, "stdout")
	pipe.WriteCloser = lo.Must(NewPipe(context.TODO(), cfg.StdoutArn))

	return pipe
}

type OutputPipe struct {
	io.WriteCloser
	Config Config
	Arn    string
	Url    string
}

func (o *OutputPipe) Close() error {
	if o.Config.Aggregator {
		err := syscall.Kill(*o.Config.Pid, 9)
		if err != nil {
			L.Error.Println("error sending sigkill to:", o.Config.Pid)
		}
		L.Debug.Println("sent sigkill to:", o.Config.Pid)
	}

	return nil
}

func NewPipe(ctx context.Context, url string) (io.ReadWriteCloser, error) {
	client := sqs.NewFromConfig(lo.Must(config.LoadDefaultConfig(ctx)))

	p := &Pipe{
		url:      url,
		client:   client,
		ctx:      ctx,
		buf:      bytes.NewBuffer([]byte{}),
		sent:     aws.Uint64(0),
		received: aws.Uint64(0),
		expected: aws.Uint64(0),
	}
	return p, nil
}

type EventType int

const (
	ShutdownEvent EventType = iota
	MessageEvent
	FilterEvent
)

type Event struct {
	Type    EventType
	Id      uint64
	Content string
}

type Pipe struct {
	ctx      context.Context
	buf      *bytes.Buffer
	closed   bool
	url      string
	client   *sqs.Client
	sent     *uint64
	received *uint64
	expected *uint64
}

func (p Pipe) Read(d []byte) (n int, err error) {
	L.Debug.Printf("reading from pipe: expected %d, received: %d\n", *p.expected, *p.received)

	for p.buf.Len() == 0 {
		if *p.expected != 0 && *p.received >= *p.expected {
			L.Debug.Println("finished processing all messages")
			return 0, io.EOF
		}

		p.closed = p.fetch()
	}

	return p.buf.Read(d)
}

func (p Pipe) fetch() bool {
	L.Debug.Println("fetching messages")
	msgs := lo.Must(p.client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
		QueueUrl:        aws.String(p.url),
		WaitTimeSeconds: 2,
	}))

	for _, msg := range msgs.Messages {
		lo.Must(p.client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
			QueueUrl:      aws.String(p.url),
			ReceiptHandle: msg.ReceiptHandle,
		}))

		L.Debug.Printf("received message: %s\n", *msg.Body)

		event := &Event{}
		err := json.Unmarshal([]byte(*msg.Body), event)
		if err != nil {
			L.Error.Fatalln("failed to unmarshal message:", err)
		}

		L.Debug.Printf("event: id %d", event.Id)
		if event.Type == ShutdownEvent {
			L.Debug.Printf("received shutdown event: expected %d", event.Id)
			*p.expected = event.Id
		} else if event.Type == MessageEvent {
			*p.received += uint64(1)

			L.Debug.Printf("received message event: id %d", event.Id)
			p.buf.Truncate(0)
			p.buf.WriteString(event.Content)
		} else {
			L.Error.Fatalf("unknown event type: %d", event.Type)
		}
	}

	return false
}

func (p Pipe) Write(d []byte) (n int, err error) {
	L.Debug.Printf("writing to pipe: %s: %s", p.url, string(d))

	msg := lo.Must(json.Marshal(&Event{
		Type:    MessageEvent,
		Id:      *p.sent + 1,
		Content: string(d),
	}))

	lo.Must(p.client.SendMessage(p.ctx, &sqs.SendMessageInput{
		QueueUrl: aws.String(p.url),
		//MessageGroupId:         aws.String("default"),
		//MessageDeduplicationId: aws.String(messageId()),
		MessageBody: aws.String(string(msg)),
	}))
	if err != nil {
		return 0, err
	}

	L.Debug.Printf("sent message: id %d: %s", *p.sent+1, string(d))

	*p.sent += 1

	return len(d), err
}

func (p Pipe) Close() (err error) {
	L.Debug.Println("closing pipes")

	event := lo.Must(json.Marshal(&Event{
		Type: ShutdownEvent,
		Id:   *p.sent,
	}))
	lo.Must(p.client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl: aws.String(p.url),
		//MessageGroupId:         aws.String("default"),
		//MessageDeduplicationId: aws.String(messageId()),
		MessageBody: aws.String(string(event)),
	}))
	L.Debug.Println("sent closed event: processed", *p.sent)

	// wait for the collection command to signal shutdown.
	lo.Must(io.Copy(os.Stdout, os.Stdin))

	return nil
}
