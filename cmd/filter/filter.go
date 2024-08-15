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
	l, err := lambda.New(app)
	if err != nil {
		return nil, fmt.Errorf("lambda: %w", err)
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
	sfn.INextable
	Start    sfn.Pass
	Lambda   *lambda.Lambda
	choice   sfn.Choice
	filtered sfn.IChainable
}

func (s Filter) GetName() string { return "filter" }

func (s *Filter) Compile(stack constructs.Construct, i int) error {
	name := fmt.Sprintf("%s-%d", s.GetName(), i)
	err := s.Lambda.Compile(stack, 0)
	if err != nil {
		return fmt.Errorf("filter: %w", err)
	}

	s.choice = sfn.NewChoice(stack, jsii.String(name+"-choice"), &sfn.ChoiceProps{
		Comment:    jsii.String("choice"),
		OutputPath: jsii.String("$.__input"),
	})

	s.Start = sfn.NewPass(stack, jsii.String(name+"-pass"), &sfn.PassProps{
		Parameters: &map[string]interface{}{
			"__input.$": "$",
		},
		Comment: jsii.String("save input to $.__input"),
	})

	s.filtered = sfn.NewSucceed(stack, jsii.String(name+"-filtered"), &sfn.SucceedProps{})
	return nil
}

func (s *Filter) Next(next sfn.IChainable) sfn.Chain {
	return s.Start.
		Next(s.Lambda).
		Next(
			s.choice.When(sfn.Condition_BooleanEquals(jsii.String("$.__choice.result"), jsii.Bool(true)),
				next, &sfn.ChoiceTransitionOptions{},
			).Otherwise(
				s.filtered,
			),
		)
}
