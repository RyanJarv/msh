package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/samber/lo"
	"io"
	"os"
	"strconv"
	"time"
)

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

func messageId() string {
	id := RandString(10) + strconv.FormatInt(time.Now().UnixNano(), 10)
	return id
}
