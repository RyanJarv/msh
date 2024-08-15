package sleep

import (
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/types"
	"os"
	"strconv"
	"syscall"
)

func New(app app.App) (*SleepCmd, error) {
	if app.NArg() != 2 {
		return nil, fmt.Errorf("usage: %s <seconds>", os.Args[0])
	}

	seconds, err := strconv.Atoi(app.Arg(1))
	if err != nil {
		return nil, fmt.Errorf("failed to parse seconds: %w", err)
	}

	return &SleepCmd{
		Seconds: seconds,
	}, nil
}

type SleepCmd struct {
	Seconds int
	types.IChain
}

func (s SleepCmd) GetName() string { return "sleep" }

func (s *SleepCmd) Compile(stack constructs.Construct, i int) error {
	name := fmt.Sprintf("%s-%d", s.GetName(), syscall.Getpid())

	s.IChain = sfn.NewWait(stack, jsii.String(name), &sfn.WaitProps{
		Time: sfn.WaitTime_Duration(awscdk.Duration_Seconds(jsii.Number(s.Seconds))),
	})

	return nil
}
