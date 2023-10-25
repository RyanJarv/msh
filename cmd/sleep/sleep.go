package sleep

import (
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/jsii-runtime-go"
	"log"
	"os"
	"strconv"
)

func NewSleep(args []string) (*SleepCmd, error) {
	if len(args) != 1 {
		log.Fatalf("usage: %s <seconds>", os.Args[0])
	}

	seconds, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse seconds: %w", err)
	}

	return &SleepCmd{
		Seconds: seconds,
	}, nil
}

type SleepCmd struct {
	Seconds int
}

func (s *SleepCmd) Name() string { return "sleep" }

func (s *SleepCmd) Compile(stack awscdk.Stack, next interface{}) (interface{}, error) {
	chain, ok := next.(awsstepfunctions.IChainable)
	if !ok {
		return nil, fmt.Errorf("next step must be statemachine chain")
	}

	wait := sfn.NewWait(stack, jsii.String("wait"), &sfn.WaitProps{
		Time: sfn.WaitTime_Duration(awscdk.Duration_Seconds(jsii.Number(s.Seconds))),
	})

	return wait.Next(chain), nil
}
