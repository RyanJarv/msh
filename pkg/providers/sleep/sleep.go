package sleep

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/jsii-runtime-go"
	"log"
	"os"
)

func NewSleep(args []string) (*SleepCmd, error) {
	if len(args) != 1 {
		log.Fatalf("usage: %s <seconds>", os.Args[0])
	}

	return &SleepCmd{
		Seconds: args[0],
	}, nil
}

type SleepCmd struct {
	Seconds string
}

func (s *SleepCmd) Name() string { return "sleep" }

func (s *SleepCmd) SfnHook(stack awscdk.Stack, chain awsstepfunctions.Chain) awsstepfunctions.Chain {
	return chain.Next(
		sfn.NewWait(stack, jsii.String("wait"), &sfn.WaitProps{
			Time: sfn.WaitTime_SecondsPath(jsii.String(s.Seconds)),
		}),
	)
}
