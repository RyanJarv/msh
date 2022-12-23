package eventbridge

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/pipes"
	pipesTypes "github.com/aws/aws-sdk-go-v2/service/pipes/types"
	"github.com/ryanjarv/msh/pkg/fd"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/providers/lambda"
	"github.com/samber/lo"
	"io"
)

func NewPipe(lambda *lambda.LambdaCmd) AwsPipe {
	cfg := lo.Must(config.LoadDefaultConfig(context.TODO()))

	name := aws.String("test2")
	return AwsPipe{
		Name: name,
		CreatePipeInput: &pipes.CreatePipeInput{
			Name:         name,
			Enrichment:   lambda.Arn(),
			DesiredState: pipesTypes.RequestedPipeStateRunning,
			RoleArn:      aws.String("arn:aws:iam::336983520827:role/service-role/Amazon_EventBridge_Pipe_test_ed80013d"),
		},
		LambdaCmd: lambda,
		cfg:       lo.Must(config.LoadDefaultConfig(context.TODO())),
		client:    pipes.NewFromConfig(cfg),
	}
}

type AwsPipe struct {
	CreatePipeInput *pipes.CreatePipeInput
	cfg             aws.Config
	LambdaCmd       *lambda.LambdaCmd
	client          *pipes.Client
	Name            *string
}

func (p *AwsPipe) Deploy() error {
	_, err := p.client.CreatePipe(context.TODO(), p.CreatePipeInput)
	t := &pipesTypes.ConflictException{}

	if errors.As(err, &t) {
		L.Debug.Println("reusing existing pipe:", p.Name)
	} else if err != nil {
		return fmt.Errorf("failed to create pipe: %w\n", err)
	}

	return nil
}

func (p *AwsPipe) Run() error {
	p.LambdaCmd.Stdout.(*fd.Sqs).Wait()
	return nil
}

func (p *AwsPipe) SetStdin(source interface{}) {
	sqs := lo.Must(fd.NewSqsFrom(context.TODO(), source, *p.Name, "stdin"))

	p.CreatePipeInput.Source = sqs.Arn()
	p.CreatePipeInput.SourceParameters = sqs.PipeSourceParameters()

	L.Debug.Println("eventbridge pipes: source:", sqs.Arn())

	p.LambdaCmd.SetStdin(sqs)
}

func (p *AwsPipe) GetStdout() io.Reader {
	stdout := p.LambdaCmd.GetStdout()

	sqs := lo.Must(fd.NewSqsFrom(context.TODO(), stdout, *p.Name, "stdout"))

	L.Debug.Println("eventbridge pipes: target:", sqs.Arn())

	p.CreatePipeInput.Target = sqs.Arn()
	p.LambdaCmd.Stdout = sqs

	return sqs
}
