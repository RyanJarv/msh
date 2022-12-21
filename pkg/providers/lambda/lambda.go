package lambda

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
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
)

// Run takes the invocation args and transparently runs the intended command while intercepting sensitive info.
func Run(args []string) error {

	name := utils.GetName(args) + "-2"
	L.Info.Println("name:", name)

	awsCfg := lo.Must(config.LoadDefaultConfig(context.TODO()))
	cfg := fd.NewConfig(awsCfg, name)
	stdout := fd.NewRemoteOutputPipe(cfg)
	fd.WriteConf(*cfg, stdout)

	stdin := fd.NewInputPipe(cfg, true)

	if cfg.WrapperConfig.StdoutLocal {
		stdin.Arn, stdin.Url = fd.SetupSqsFd(cfg.AwsCfg, cfg.Name, "stdin")
		L.Debug.Println("forwarding: stdin ->", stdin.Arn)
	}

	remoteStdin := lo.Must(fd.NewPipe(context.TODO(), stdin.Url))
	//setupPipes(awsCfg, name, "arn:aws:lambda:us-east-1:336983520827:function:msh-test", stdin, stdout)

	fmt.Println("13")
	if cfg.WrapperConfig.StdoutLocal {
		lo.Must(io.Copy(remoteStdin, stdin))
		return nil
	} else {
		stdout.Close()
		io.Copy(os.Stdout, os.Stdin)
		return nil
	}
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

func setupPipes(cfg aws.Config, name, lambdaArn string, stdin *fd.InputPipe, stdout *fd.OutputPipe) {

	client := pipes.NewFromConfig(cfg)

	fmt.Println("stdin:", stdin.Arn)
	fmt.Println("stdout:", stdout.Arn)

	_, err := client.CreatePipe(
		context.TODO(),
		&pipes.CreatePipeInput{
			Name:         aws.String(name),
			RoleArn:      aws.String("arn:aws:iam::336983520827:role/service-role/Amazon_EventBridge_Pipe_test_ed80013d"),
			Source:       aws.String(stdin.Arn),
			Target:       aws.String(stdout.Arn),
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
