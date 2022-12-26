package fd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	pipesTypes "github.com/aws/aws-sdk-go-v2/service/pipes/types"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/types"
	"github.com/ryanjarv/msh/pkg/utils"
	"github.com/samber/lo"
	"io"
	"strings"
	"sync"
)

type EventType int

const (
	ShutdownEvent EventType = iota
	MessageEvent

	// Not implemented
	//FilterEvent = 2
)

type Referable interface {
	Arn() *string
	Url() *string
}

type ISqs interface {
	io.ReadWriteCloser
	Referable
	Wait()
	PipeSourceParameters() *pipesTypes.PipeSourceParameters
	WritePolicy() *types.IamPolicy
	ReadPolicy() *types.IamPolicy
}

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

func CreateSqs(cfg aws.Config, name, fd string) (ISqs, error) {
	_, url := SetupSqsFd(cfg, name, fd)
	client := sqs.NewFromConfig(lo.Must(config.LoadDefaultConfig(context.TODO())))

	p := &Sqs{
		SqsUrl:   &url,
		client:   client,
		ctx:      context.Background(),
		fd:       fd,
		name:     fmt.Sprintf("%s-%s", name, fd),
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
		SqsUrl:   &url,
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

func NewSqsFrom(ctx context.Context, f interface{}, name, fd string) (ISqs, error) {
	var pipe ISqs

	switch from := f.(type) {
	case ISqs:
		// Just read directly from the sqs output queue of the last command.
		L.Debug.Printf("%s is sqs, skipping copy: %s\n", fd, *from.Url())
		pipe = lo.Must(OpenSqs(ctx, *from.Url()))
	case io.Reader:
		cfg := lo.Must(config.LoadDefaultConfig(ctx))
		pipe = lo.Must(CreateSqs(cfg, name, fd))

		L.Debug.Printf("copying: %s (%T) -> %s\n", fd, f, *pipe.Url())

		go MustCopy(pipe, from)
	default:
		panic("not implemented")
	}

	return pipe, nil
}

type Event struct {
	Type    EventType
	Id      uint64
	Content string
	Stderr  string
}

func (e *Event) MarshalJSON() ([]byte, error) {
	return utils.JSONMarshal(map[string]interface{}{
		"Type":    e.Type,
		"Id":      e.Id,
		"Content": e.Content,
	})
}

type Sqs struct {
	ctx      context.Context
	buf      *bytes.Buffer
	wg       *sync.WaitGroup
	closed   bool
	SqsUrl   *string
	client   *sqs.Client
	sent     *uint64
	received *uint64
	expected *uint64
	fd       string
	attrs    map[string]interface{}
	name     string
}

func (s *Sqs) Read(d []byte) (n int, err error) {
	L.Debug.Printf("reading from pipe: expected %d, received: %d\n", *s.expected, *s.received)
	if s.closed {
		L.Debug.Println("pipe is closed, returning EOF")
		return 0, io.EOF
	}

	for s.buf.Len() == 0 {
		if *s.expected != 0 && *s.received >= *s.expected {
			L.Debug.Println("finished processing all messages: expected", *s.expected, "received", *s.received)

			s.wg.Done()
			s.closed = true

			return 0, io.EOF
		}

		s.fetch()
	}

	L.Debug.Println("buffer length:", s.buf.Len())

	return s.buf.Read(d)
}

func (s *Sqs) fetch() bool {
	L.Debug.Println("fetching messages")
	msgs := lo.Must(s.client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
		QueueUrl:        s.SqsUrl,
		WaitTimeSeconds: 2,
	}))

	for _, msg := range msgs.Messages {
		lo.Must(s.client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
			QueueUrl:      s.SqsUrl,
			ReceiptHandle: msg.ReceiptHandle,
		}))

		// Msg is a json string so unmarshal twice.
		var body string
		err := json.Unmarshal([]byte(*msg.Body), &body)
		if err != nil {
			L.Error.Fatalln("failed to unmarshal message body:", err)
		}

		L.Debug.Printf("received message (decoded): %s\n", body)

		event := &Event{}
		err = json.Unmarshal([]byte(body), event)
		if err != nil {
			L.Error.Fatalln("unmarshal message to event:", err)
		}

		if event.Stderr != "" {
			L.Info.Println("stderr:", event.Stderr)
		}

		L.Debug.Printf("event: id %d", event.Id)
		if event.Type == ShutdownEvent {
			L.Debug.Printf("received shutdown event: expected %d", event.Id)
			*s.expected = event.Id
		} else if event.Type == MessageEvent {
			*s.received += uint64(1)

			L.Debug.Printf("received message event: id %d", event.Id)
			s.buf.Truncate(0)
			s.buf.WriteString(event.Content)
		} else {
			L.Error.Fatalf("unknown event type: %d", event.Type)
		}
	}

	return false
}

func (s *Sqs) Arn() *string {
	return aws.String(QueueArnFromUrl(*s.SqsUrl))
}

func (s *Sqs) Url() *string {
	return s.SqsUrl
}

func (s *Sqs) Env() []string {
	return []string{
		fmt.Sprintf("AWS_SQS_%s_URL=%s", strings.ToUpper(s.fd), *s.SqsUrl),
	}
}

// Write writes the contents of p to the SQS queue.
//
// I'm unable to get eventbridge pipes w/ a transform to work when the input is json. I'm not sure if this is a bug in
// eventbridge or if I'm doing something wrong. To work around this, marshal the json twice to make it a string of json.
func (s *Sqs) Write(d []byte) (n int, err error) {
	L.Debug.Printf("writing to pipe: %s: %s", *s.SqsUrl, string(d))

	msg := event(&Event{
		Type:    MessageEvent,
		Id:      *s.sent + 1,
		Content: string(d),
	})

	lo.Must(s.client.SendMessage(s.ctx, &sqs.SendMessageInput{
		QueueUrl:    s.SqsUrl,
		MessageBody: aws.String(msg),
	}))
	if err != nil {
		return 0, err
	}

	L.Debug.Printf("sent message: id %d: %s", *s.sent+1, string(d))

	*s.sent += 1

	return len(d), err
}

func (s *Sqs) Wait() {
	s.wg.Wait()
}

func (s *Sqs) Close() (err error) {
	L.Debug.Println("closing pipes")

	body := event(&Event{
		Type: ShutdownEvent,
		Id:   *s.sent,
	})

	lo.Must(s.client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    s.SqsUrl,
		MessageBody: aws.String(body),
	}))
	L.Debug.Println("sent closed event: processed", *s.sent)

	return nil
}

func event(event *Event) string {
	// Ensure msg is a json string so marshal twice.
	msg := lo.Must(utils.JSONMarshal(event))
	msg = lo.Must(utils.JSONMarshal(string(msg)))

	L.Debug.Println("msg:", string(msg))

	return string(msg)
}

func (s *Sqs) PipeSourceParameters() *pipesTypes.PipeSourceParameters {
	return &pipesTypes.PipeSourceParameters{
		SqsQueueParameters: &pipesTypes.PipeSourceSqsQueueParameters{
			BatchSize:                      aws.Int32(1),
			MaximumBatchingWindowInSeconds: aws.Int32(10),
		},
	}
}

func (s *Sqs) WritePolicy() *types.IamPolicy {
	return &types.IamPolicy{
		Version: "2012-10-17",
		Name:    fmt.Sprintf("%s-write", s.name),
		Statement: []types.IamPolicyStatement{
			{
				Effect: "Allow",
				Action: []string{
					"sqs:SendMessage",
				},
				Resource: []string{
					*s.Arn(),
				},
			},
		},
	}
}

func (s *Sqs) ReadPolicy() *types.IamPolicy {
	return &types.IamPolicy{
		Version: "2012-10-17",
		Name:    fmt.Sprintf("%s-read", s.name),
		Statement: []types.IamPolicyStatement{
			{
				Effect: "Allow",
				Action: []string{
					"sqs:ReceiveMessage",
					"sqs:DeleteMessage",
					"sqs:GetQueueAttributes",
				},
				Resource: []string{
					*s.Arn(),
				},
			},
		},
	}
}
