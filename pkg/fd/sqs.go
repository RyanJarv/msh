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
	"sync"
)

////type Processor interface {
////	Setup(input Input, output Output)
////	Run() error
////}
////
////type Process struct {
////	Processor Processor
////}
////
////func setupFds(cfg aws.Config, name string) (io.ReadCloser, io.WriteCloser, func()) {
////	var stdin io.ReadCloser
////	var pid int
////	if utils.IsTTY(os.Stdin) {
////		L.Debug.Println("stdin: terminal")
////		stdin = os.Stdin
////	} else {
////		conf := ReadConf(os.Stdin)
////		stdinArn = conf.Url
////		pid = *conf.Pid
////
////		L.Debug.Println("stdin:", conf.Url)
////		stdin = lo.Must(utils.OpenSqs(context.TODO(), conf.Url))
////	}
////
////	stdoutArn, stdoutUrl := SetupSqsFd(cfg, name, "stdout")
////
////	stdout := lo.Must(utils.OpenSqs(context.TODO(), stdoutUrl))
////
////	var read func()
////	if utils.IsTTY(os.Stdout) {
////		L.Debug.Println("stdout: terminal")
////		stdout = os.Stdout
////
////		L.Debug.Printf("stdout: %s -> terminal\n", stdoutArn)
////
////		read = func() {
////			lo.Must(io.Copy(os.Stdout, stdout))
////			err := syscall.Kill(pid, 9)
////			if err != nil {
////				L.Error.Println("sending sigkill to:", pid)
////			}
////			L.Debug.Println("sent sigkill to:", pid)
////		}
////	} else {
////		arn, Url := SetupSqsFd(cfg, name, "stdout")
////		L.Debug.Println("stdout:", arn)
////		stdout = lo.Must(utils.OpenSqs(context.TODO(), Url))
////
////		WriteConf(os.Stdout, WrapperConfig{
////			Url: Url,
////			Arn: arn,
////			Pid:       aws.Int(os.Getpid()),
////		})
////	}
////
////	return stdin, stdout, read
////}
//

//
//func NewInputPipe(cfg *conf.Config, remote bool) *InputPipe {
//	pipe := &InputPipe{}
//	if cfg.WrapperConfig.StdoutLocal {
//		L.Debug.Println("stdin: terminal")
//		pipe.ReadCloser = os.Stdin
//	} else {
//		L.Debug.Println("stdin:", cfg.StdoutUrl)
//		pipe.Url = cfg.WrapperConfig.StdoutUrl
//		pipe.Arn = cfg.WrapperConfig.StdoutArn
//		pipe.ReadCloser = lo.Must(OpenSqs(context.TODO(), cfg.WrapperConfig.StdoutUrl))
//	}
//
//	return pipe
//}
//
//type InputPipe struct {
//	io.ReadCloser
//	Arn string
//	Url string
//	pid conf.Config
//}
//
//func NewRemoteOutputPipe(cfg *conf.Config) *OutputPipe {
//	pipe := &OutputPipe{}
//
//	pipe.Arn, pipe.Url = SetupSqsFd(cfg.AwsCfg, cfg.Name, "stdout")
//	pipe.WriteCloser = lo.Must(OpenSqs(context.TODO(), cfg.StdoutArn))
//
//	return pipe
//}
//
//type OutputPipe struct {
//	io.WriteCloser
//	Config conf.Config
//	Arn    string
//	Url    string
//}
//
//func (o *OutputPipe) Close() error {
//	if o.Config.Aggregator {
//		err := syscall.Kill(*o.Config.Pid, 9)
//		if err != nil {
//			L.Error.Println("error sending sigkill to:", o.Config.Pid)
//		}
//		L.Debug.Println("sent sigkill to:", o.Config.Pid)
//	}
//
//	return nil
//}

func QueueArnFromUrl(url string) string {
	// https://sqs.us-east-1.amazonaws.com/123456789012/cat-msh-0a0c3f9d7adb3006a923c4fd8884951b068f44b5-stdout3
	p := strings.Split(url, "/")
	domain := p[2]
	region := strings.Split(domain, ".")[1]

	account := p[3]
	name := p[4]

	return fmt.Sprintf("arn:aws:sqs:%s:%s:%s", region, account, name)
}

func SetupSqsFd(cfg aws.Config, name, fd string) (arn string, uri string) {
	L.Debug.Println("creating sqs queue:", name, fd)
	client := sqs.NewFromConfig(cfg)

	queueName := fmt.Sprintf("%s-%s", name, fd)
	resp := lo.Must(client.CreateQueue(context.TODO(), &sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
	}))
	L.Debug.Println("created sqs queue:", *resp.QueueUrl)

	return QueueArnFromUrl(*resp.QueueUrl), *resp.QueueUrl
}

func CreateSqs(cfg aws.Config, name, fd string) (*Sqs, error) {
	_, url := SetupSqsFd(cfg, name, fd)
	client := sqs.NewFromConfig(lo.Must(config.LoadDefaultConfig(context.TODO())))

	p := &Sqs{
		Url:      &url,
		client:   client,
		ctx:      context.Background(),
		fd:       fd,
		wg:       &sync.WaitGroup{},
		buf:      bytes.NewBuffer([]byte{}),
		sent:     aws.Uint64(0),
		received: aws.Uint64(0),
		expected: aws.Uint64(0),
	}
	p.wg.Add(1)
	return p, nil
}

func OpenSqs(ctx context.Context, url string) (*Sqs, error) {
	client := sqs.NewFromConfig(lo.Must(config.LoadDefaultConfig(ctx)))

	p := &Sqs{
		Url:      &url,
		client:   client,
		ctx:      ctx,
		wg:       &sync.WaitGroup{},
		buf:      bytes.NewBuffer([]byte{}),
		sent:     aws.Uint64(0),
		received: aws.Uint64(0),
		expected: aws.Uint64(0),
	}

	p.wg.Add(1)
	return p, nil
}

func NewSqsFrom(ctx context.Context, f interface{}) (*Sqs, error) {
	var pipe *Sqs
	switch from := f.(type) {

	// Just read directly from the sqs output queue of the last command.
	case map[string]interface{}:
		url := from["Url"].(string)
		L.Debug.Println("stdin is sqs, skipping copy:", url)
		pipe = lo.Must(OpenSqs(ctx, url))
	case Sqs:
		// Just read directly from the sqs output queue of the last command.
		L.Debug.Println("stdin is sqs, skipping copy:", from.Url)
		pipe = lo.Must(OpenSqs(ctx, *from.Url))
	case io.Reader:
		cfg := lo.Must(config.LoadDefaultConfig(ctx))
		pipe = lo.Must(CreateSqs(cfg, "temp", "stdin"))

		L.Debug.Printf("copying to:  %s", *pipe.Url)
		go MustCopy(pipe, from)
	default:
		panic("not implemented")
	}

	return pipe, nil
}

type EventType int

const (
	ShutdownEvent EventType = iota
	MessageEvent
	//FilterEvent
)

type Event struct {
	Type    EventType
	Id      uint64
	Content string
}

type Sqs struct {
	io.ReadWriteCloser
	Referable
	ctx      context.Context
	buf      *bytes.Buffer
	wg       *sync.WaitGroup
	closed   bool
	Url      *string
	client   *sqs.Client
	sent     *uint64
	received *uint64
	expected *uint64
	fd       string
}

func (p Sqs) Read(d []byte) (n int, err error) {
	L.Debug.Printf("reading from pipe: expected %d, received: %d\n", *p.expected, *p.received)

	for p.buf.Len() == 0 {
		if *p.expected != 0 && *p.received >= *p.expected {
			L.Debug.Println("finished processing all messages")
			return 0, io.EOF
		}

		if p.fetch() {
			p.wg.Done()
		}
	}

	return p.buf.Read(d)
}

func (p Sqs) fetch() bool {
	L.Debug.Println("fetching messages")
	msgs := lo.Must(p.client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
		QueueUrl:        p.Url,
		WaitTimeSeconds: 2,
	}))

	for _, msg := range msgs.Messages {
		lo.Must(p.client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
			QueueUrl:      p.Url,
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

func (p Sqs) Arn() *string {
	return aws.String(QueueArnFromUrl(*p.Url))
}

func (p Sqs) Env() []string {
	return []string{
		fmt.Sprintf("AWS_SQS_%s_URL=%s", strings.ToUpper(p.fd), *p.Url),
	}
}

func (p Sqs) Write(d []byte) (n int, err error) {
	L.Debug.Printf("writing to pipe: %s: %s", *p.Url, string(d))

	msg := lo.Must(json.Marshal(&Event{
		Type:    MessageEvent,
		Id:      *p.sent + 1,
		Content: string(d),
	}))

	lo.Must(p.client.SendMessage(p.ctx, &sqs.SendMessageInput{
		QueueUrl: p.Url,
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

func (p Sqs) Wait() {
	p.wg.Wait()
}

func (p Sqs) Close() (err error) {
	L.Debug.Println("closing pipes")

	event := lo.Must(json.Marshal(&Event{
		Type: ShutdownEvent,
		Id:   *p.sent,
	}))
	lo.Must(p.client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl: p.Url,
		//MessageGroupId:         aws.String("default"),
		//MessageDeduplicationId: aws.String(messageId()),
		MessageBody: aws.String(string(event)),
	}))
	L.Debug.Println("sent closed event: processed", *p.sent)

	// wait for the collection command to signal shutdown.
	lo.Must(io.Copy(os.Stdout, os.Stdin))

	return nil
}
