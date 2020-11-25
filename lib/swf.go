package lib

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	L "github.com/ryanjarv/msh/logger"
)

func Swf(argv []string) (err error) {
	path := argv[0]

	var json string
	if json, err = ReadConfig(path); err != nil {
		return err
	}

	var client *sfn.Client
	if cfg, err := config.LoadDefaultConfig(); err != nil {
		return err
	} else {
		client = sfn.NewFromConfig(cfg)
	}

	ctx := context.Background()
	machine, err := client.CreateStateMachine(ctx, &sfn.CreateStateMachineInput{
		Definition: aws.String(json),
		Name:       aws.String(CleanName(path)),
		RoleArn:    aws.String("arn:aws:iam::253528964770:role/config-swf-hello-world-iam-StatesRole-28JGD3HEQ0Q5"), //TODO: unhardcode this
		LoggingConfiguration: &types.LoggingConfiguration{
			Destinations: []*types.LogDestination{
				{
					CloudWatchLogsLogGroup: &types.CloudWatchLogsLogGroup{
						LogGroupArn: aws.String("arn:aws:logs:us-east-2:253528964770:log-group:/aws/vendedlogs/states/Helloworld-Logs:*"), //TODO: unhardcode
					},
				},
			},
			IncludeExecutionData: aws.Bool(true),
			Level:                "ALL",
		},
		Tags:                 nil,
		TracingConfiguration: nil,
		Type:                 "",
	})
	if err == nil {
		L.Debug.Printf("State swf: %v\n", *machine)
	}
	return err
}

