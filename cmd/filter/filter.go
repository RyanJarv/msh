package filter

import (
	_ "embed"
	"fmt"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/cmd/lambda"
	"github.com/ryanjarv/msh/pkg/app"
)

func New(app app.App) (*Filter, error) {
	return &Filter{
		Args: app.Args(),
		Opts: &lambda.LambdaOpts{
			InputPath:      jsii.String("$.__input"),
			ResultSelector: &map[string]interface{}{"result.$": "$.Payload"},
			ResultPath:     jsii.String("$.__choice"),
		},
	}, nil
}

type Filter struct {
	sfn.INextable `json:"-"`
	start         sfn.Pass
	choice        sfn.Choice
	filtered      sfn.IChainable
	lambda        *lambda.Lambda
	Args          []string
	Opts          *lambda.LambdaOpts
}

func (s Filter) GetName() string { return "filter" }

func (s *Filter) Compile(stack constructs.Construct, i int) error {
	name := fmt.Sprintf("%s-%d", s.GetName(), i)

	var err error
	s.lambda, err = lambda.New(s.Args)
	if err != nil {
		return fmt.Errorf("%s: compile: %w", name, err)
	}

	s.lambda.SetOpts(s.Opts)

	if err = s.lambda.Compile(stack, 0); err != nil {
		return fmt.Errorf("filter: %w", err)
	}

	s.choice = sfn.NewChoice(stack, jsii.String(name+"-choice"), &sfn.ChoiceProps{
		Comment:    jsii.String("choice"),
		OutputPath: jsii.String("$.__input"),
	})

	s.start = sfn.NewPass(stack, jsii.String(name+"-pass"), &sfn.PassProps{
		Parameters: &map[string]interface{}{
			"__input.$": "$",
		},
		Comment: jsii.String("save input to $.__input"),
	})

	s.filtered = sfn.NewSucceed(stack, jsii.String(name+"-filtered"), &sfn.SucceedProps{})
	return nil
}

func (s *Filter) Next(next sfn.IChainable) sfn.Chain {
	return s.start.
		Next(s.lambda).
		Next(
			s.choice.When(sfn.Condition_BooleanEquals(jsii.String("$.__choice.result"), jsii.Bool(true)),
				next, &sfn.ChoiceTransitionOptions{},
			).Otherwise(
				s.filtered,
			),
		)
}
