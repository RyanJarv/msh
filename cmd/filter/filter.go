package filter

import (
	_ "embed"
	"fmt"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/cmd/lambda"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/types"
)

func New(app app.App) (*Cmd, error) {
	return &Cmd{
		Args: app.Args()[1:],
		Opts: &lambda.LambdaOpts{
			InputPath:      jsii.String("$.__input"),
			ResultSelector: &map[string]interface{}{"result.$": "$.Payload"},
			ResultPath:     jsii.String("$.__choice"),
		},
	}, nil
}

type Cmd struct {
	types.IChain `json:"-"`
	Args         []string
	Opts         *lambda.LambdaOpts
	stack        constructs.Construct
	name         string
}

func (s Cmd) GetName() string { return "filter" }

func (s *Cmd) Compile(stack constructs.Construct, i int) error {
	s.name = fmt.Sprintf("%s-%d", s.GetName(), i)
	s.stack = stack

	lambda, err := lambda.NewInternal(s.Args)
	if err != nil {
		return fmt.Errorf("%s: %w", s.name, err)
	}

	lambda.SetOpts(s.Opts)

	if err = lambda.Compile(stack, 0); err != nil {
		return fmt.Errorf("filter: %w", err)
	}

	s.IChain = lambda

	return nil
}

func (s *Cmd) Next(next sfn.IChainable) sfn.Chain {
	return sfn.NewPass(s.stack, jsii.String(s.name+"-pass"), &sfn.PassProps{
		Parameters: &map[string]interface{}{
			"__input.$": "$",
		},
		Comment: jsii.String("save input to $.__input"),
	}).
		Next(s.IChain.StartState()).
		Next(
			sfn.NewChoice(s.stack, jsii.String(s.name+"-choice"), &sfn.ChoiceProps{
				Comment:    jsii.String("choice"),
				OutputPath: jsii.String("$.__input"),
			}).When(sfn.Condition_BooleanEquals(jsii.String("$.__choice.result"), jsii.Bool(true)),
				next.StartState(), &sfn.ChoiceTransitionOptions{},
			).Otherwise(sfn.NewSucceed(s.stack, jsii.String(s.name+"-filtered"), &sfn.SucceedProps{}).StartState()),
		)
}
