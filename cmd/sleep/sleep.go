package sleep

import (
	"flag"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/types"
	"github.com/ryanjarv/msh/pkg/utils"
	"os"
	"strconv"
)

func New(app *app.App, argv []string) (types.CdkStep, error) {
	flags := flag.NewFlagSet("sleep", flag.ExitOnError)
	flags.Parse(argv)

	if flags.NArg() != 2 {
		return nil, fmt.Errorf("usage: %s <seconds>", os.Args[0])
	}

	seconds, err := strconv.Atoi(flags.Arg(1))
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

func (s *SleepCmd) GetName() string { return "sleep" }

func (s *SleepCmd) GetChain(step *types.StepRunInfo) sfn.IChainable {
	wait := sfn.NewWait(step.Scope, utils.GetId(s.GetName(), step.Index), &sfn.WaitProps{
		Time: sfn.WaitTime_Duration(awscdk.Duration_Seconds(jsii.Number(s.Seconds))),
	})
	if step.Next == nil {
		return wait
	} else if v, ok := step.Next.Step.(types.SfnChainable); ok {
		return sfn.Chain_Sequence(wait, v.GetChain(step.Next))
	}

	return nil
}
