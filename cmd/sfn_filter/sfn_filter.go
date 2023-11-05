package sfn_filter

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
	l, err := lambda.New(app)
	if err != nil {
		return nil, err
	}

	l.SetOpts(&lambda.LambdaOpts{
		InputPath:      jsii.String("$.__input"),
		ResultSelector: &map[string]interface{}{"result.$": "$.Payload"},
		ResultPath:     jsii.String("$.__choice"),
	})
	return &Filter{
		Lambda: l,
	}, nil
}

type Filter struct {
	*lambda.Lambda
	sfn.Pass
	choice   sfn.Choice
	filtered sfn.IChainable
}

func (s Filter) GetName() string { return "filter" }

func (s *Filter) Compile(stack constructs.Construct, i int) error {
	err := s.Lambda.Compile(stack, 0)
	if err != nil {
		return fmt.Errorf("filter: %w", err)
	}

	s.choice = sfn.NewChoice(stack, jsii.String("choice"), &sfn.ChoiceProps{
		Comment:    jsii.String("choice"),
		OutputPath: jsii.String("$.__input"),
	})

	s.Pass = sfn.NewPass(stack, jsii.String("pass"), &sfn.PassProps{
		Parameters: &map[string]interface{}{
			"__input.$": "$",
		},
		Comment: jsii.String("save input to $.__input"),
	})

	s.filtered = sfn.NewSucceed(stack, jsii.String("filtered"), &sfn.SucceedProps{})
	return nil
}

func (s *Filter) Next(next sfn.Chain) sfn.Chain {
	return s.Pass.
		Next(s.Lambda).
		Next(
			s.choice.When(sfn.Condition_BooleanEquals(jsii.String("$.__choice.result"), jsii.Bool(true)),
				next, &sfn.ChoiceTransitionOptions{},
			).Otherwise(
				s.filtered,
			),
		)
}
