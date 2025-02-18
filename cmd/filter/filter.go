package filter

import (
	_ "embed"
	"flag"
	"fmt"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/cmd/lambda"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/types"
)

func New(_ *app.App, argv []string) (types.CdkStep, error) {
	flags := flag.NewFlagSet("filter", flag.ExitOnError)
	flags.Parse(argv[1:])

	return &Cmd{
		Args: flags.Args(),
		Opts: &lambda.LambdaOpts{
			InputPath:      jsii.String("$.__input"),
			ResultSelector: &map[string]interface{}{"result.$": "$.Payload"},
			ResultPath:     jsii.String("$.__choice"),
		},
	}, nil
}

type Cmd struct {
	*lambda.Lambda
	Args      []string
	Opts      *lambda.LambdaOpts
	RunLambda sfn.IChainable
}

func (s Cmd) GetName() string { return "filter" }

func (s *Cmd) BeforeRun(step *types.StepRunInfo) error {
	var err error
	s.Lambda, err = lambda.NewInternal(s.Args)
	if err != nil {
		return fmt.Errorf("%s: %w", s.GetName(), err)
	}

	if err := s.Lambda.BeforeRun(step); err != nil {
		return fmt.Errorf("filter: %w", err)
	}

	s.Lambda.SetOpts(s.Opts)

	return nil
}

func (s *Cmd) GetChain(i *types.StepRunInfo) sfn.IChainable {
	responseTrue := sfn.Condition_BooleanEquals(jsii.String("$.__choice.result"), jsii.Bool(true))
	filtered := sfn.NewSucceed(i.Scope, i.Id("filtered"), &sfn.SucceedProps{}).StartState()

	return sfn.NewPass(i.Scope, i.Id("pass"), &sfn.PassProps{
		Parameters: &map[string]interface{}{
			"__input.$": "$",
		},
		Comment: jsii.String("save input to $.__input"),
	}).
		Next(s.LambdaInvoke).
		Next(
			sfn.NewChoice(i.Scope, i.Id("choice"), &sfn.ChoiceProps{
				Comment:    jsii.String("choice"),
				OutputPath: jsii.String("$.__input"),
			}).When(responseTrue, i.GetChain(), &sfn.ChoiceTransitionOptions{}).
				Otherwise(filtered),
		).StartState()
}
