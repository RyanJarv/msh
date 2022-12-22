package pipes

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/pipes"
	"github.com/ryanjarv/msh/pkg/fd"
	"github.com/samber/lo"
	"io"
	"sync"
	"time"
)

func NewPipe() AwsPipe {
	return AwsPipe{
		CreatePipeInput: &pipes.CreatePipeInput{},
		cfg:             lo.Must(config.LoadDefaultConfig(context.TODO())),
	}
}

type AwsPipe struct {
	*pipes.CreatePipeInput
	cfg aws.Config
}

type Step interface {
	Arn() *string
}

func (s AwsPipe) Run() error {
	//s.CreatePipeInput.Enrichment = step.Arn()
	time.Sleep(time.Second * 1000)
	return nil
}

func (s AwsPipe) SetStdin(p interface{}) {
	sqs := lo.Must(fd.NewSqsFrom(context.TODO(), p))
	s.CreatePipeInput.Source = sqs.Arn()
}

func (s AwsPipe) GetStdout() io.Reader {
	cfg := lo.Must(config.LoadDefaultConfig(context.TODO()))
	sqs := lo.Must(fd.CreateSqs(cfg, "temp", "stdout"))
	s.CreatePipeInput.Target = sqs.Arn()
	return sqs
}

func (s AwsPipe) NewFd(name string) *fd.Sqs {
	_, uri := fd.SetupSqsFd(s.cfg, "test", name)
	return lo.Must(fd.OpenSqs(context.TODO(), uri))
}

type LocalPipe struct {
	io.ReadCloser
	Ref string
}

func Block() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
