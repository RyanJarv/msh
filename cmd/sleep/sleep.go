package sleep

import (
	"flag"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/utils"
	"os"
	"strconv"
)

func New(app app.App) (*SleepCmd, error) {
	fmt.Println(app.Args)
	flags := utils.ParseArgs(app.Args)
	if flag.NArg() != 1 {
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

func (s SleepCmd) GetName() string { return "sleep" }

func (s SleepCmd) Compile(stack constructs.Construct, next interface{}, i int) ([]interface{}, error) {
	var this awsstepfunctions.INextable

	this = sfn.NewWait(stack, jsii.String(s.GetName()), &sfn.WaitProps{
		Time: sfn.WaitTime_Duration(awscdk.Duration_Seconds(jsii.Number(s.Seconds))),
	})

	if next != nil {
		chain, ok := next.(awsstepfunctions.IChainable)
		if !ok {
			return nil, fmt.Errorf("next step must be statemachine chain")
		}

		this = this.Next(chain)
	}

	return []interface{}{this}, nil
}
