package lambda

import (
	"context"
	_ "embed"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	lambdaTypes "github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/aws-sdk-go-v2/service/pipes"
	pipesTypes "github.com/aws/aws-sdk-go-v2/service/pipes/types"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/utils"
	"github.com/ryanjarv/msh/pkg/utils/fd"
	"github.com/samber/lo"
	_ "gocloud.dev/pubsub/awssnssqs"
	"io"
	"os"
	"syscall"
)

// Run takes the invocation args and transparently runs the intended command while intercepting sensitive info.
func Run(args []string) error {
	cfg := lo.Must(config.LoadDefaultConfig(context.TODO()))

	name := utils.NameFromCmd(args) + "-2"
	L.Info.Println("name:", name)

	stdin, stdout, read := setupFds(cfg, name)

	//_ = deploy(cfg, name, args)
	//setupPipes(cfg, arn)
	setupPipes(cfg, name, "arn:aws:lambda:us-east-1:336983520827:function:msh-test", stdin, stdout)

	read()
	return nil
}

func setupFds(cfg aws.Config, name string) (stdinArn, stdoutArn string, read func()) {
	var pid int

	if utils.IsTTY(os.Stdin) {
		var stdinUrl string
		stdinArn, stdinUrl = fd.SetupSqsFd(cfg, name, "stdin")
		L.Debug.Printf("stdin: terminal -> %s\n", stdinArn)
		stdin := lo.Must(utils.NewPipe(context.TODO(), stdinUrl))
		go func() {
			io.Copy(stdin, os.Stdin)
		}()
	} else {
		conf := fd.ReadConf(os.Stdin)
		stdinArn = conf.StdinArn
		pid = *conf.Pid

		L.Debug.Println("stdin:", stdinArn)
	}

	stdoutArn, stdoutUrl := fd.SetupSqsFd(cfg, name, "stdout")

	stdout := lo.Must(utils.NewPipe(context.TODO(), stdoutUrl))

	if utils.IsTTY(os.Stdout) {
		L.Debug.Printf("stdout: %s -> terminal\n", stdoutArn)

		read = func() {
			lo.Must(io.Copy(os.Stdout, stdout))
			err := syscall.Kill(pid, 9)
			if err != nil {
				L.Error.Println("sending sigkill to:", pid)
			}
			L.Debug.Println("sent sigkill to:", pid)
		}

	} else {
		L.Debug.Println("stdout:", stdoutArn)
		fd.WriteConf(os.Stdout, fd.StateJson{StdinArn: stdoutArn, StdinUrl: stdoutUrl})
	}

	return stdinArn, stdoutArn, read
}

func deploy(cfg aws.Config, name string, args []string) string {
	client := lambda.NewFromConfig(cfg)
	resp := lo.Must(client.CreateFunction(context.TODO(), &lambda.CreateFunctionInput{
		Code: &lambdaTypes.FunctionCode{
			ImageUri: aws.String("336983520827.dkr.ecr.us-east-1.amazonaws.com/msh/msh-default:latest"),
		},
		FunctionName:  aws.String(name),
		Role:          aws.String("arn:aws:iam::336983520827:role/service-role/test-role-egq4voof"),
		Architectures: []lambdaTypes.Architecture{lambdaTypes.ArchitectureX8664},
		//ImageConfig:          nil,
		//MemorySize:           nil,
		PackageType: lambdaTypes.PackageTypeImage,
		Publish:     true,
	}))

	return *resp.FunctionArn
}

func setupPipes(cfg aws.Config, name, lambdaArn, stdin, stdout string) {

	client := pipes.NewFromConfig(cfg)

	_, err := client.CreatePipe(
		context.TODO(),
		&pipes.CreatePipeInput{
			Name:         aws.String(name),
			RoleArn:      aws.String("arn:aws:iam::336983520827:role/service-role/Amazon_EventBridge_Pipe_test_ed80013d"),
			Source:       aws.String(stdin),
			Target:       aws.String(stdout),
			DesiredState: pipesTypes.RequestedPipeStateRunning,
			SourceParameters: &pipesTypes.PipeSourceParameters{
				SqsQueueParameters: &pipesTypes.PipeSourceSqsQueueParameters{
					BatchSize:                      aws.Int32(1),
					MaximumBatchingWindowInSeconds: nil,
				},
			},
			Enrichment: aws.String(lambdaArn),
			//EnrichmentParameters: &pipesTypes.PipeEnrichmentParameters{
			//	InputTemplate: aws.String(`{
			//	  "pipeName": <aws.pipes.pipe-name>,
			//	  "msg": <$.body>,
			//	  "event": <$.attributes.Event>,
			//      "awsRegion": <$.awsRegion>,
			//	  "originalSqsEvent": <aws.pipes.event>
			//	}`),
			//},
			TargetParameters: &pipesTypes.PipeTargetParameters{
				SqsQueueParameters: &pipesTypes.PipeTargetSqsQueueParameters{
					//MessageDeduplicationId: aws.String("$.attributes.MessageDeduplicationId"),
					//MessageGroupId:         aws.String("default"),
				},
			},
			Tags: nil,
		},
	)
	if err != nil {
		t := &pipesTypes.ConflictException{}
		if errors.As(err, &t) {
			L.Debug.Println("reusing existing pipe:", name)
		} else {
			L.Error.Fatalln("failed to create pipe: %e\n", err)
		}
	}
}
