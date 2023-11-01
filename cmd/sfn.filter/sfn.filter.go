package sfn_filter

import (
	_ "embed"
	"fmt"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/cmd/lambda"
)

func New(args []string) (*Filter, error) {
	lambda, err := lambda.New(args, &lambda.LambdaOpts{
		InputPath:      jsii.String("$.__input"),
		ResultSelector: &map[string]interface{}{"result.$": "$.Payload"},
		ResultPath:     jsii.String("$.__choice"),
	})
	if err != nil {
		return nil, err
	}

	return &Filter{
		Lambda: lambda,
	}, nil
}

type Filter struct {
	Lambda *lambda.Lambda
}

func (s Filter) GetName() string { return "filter" }

func (s Filter) Compile(stack constructs.Construct, _next interface{}, i int) ([]interface{}, error) {
	next, ok := _next.(sfn.IChainable)
	if !ok {
		return nil, fmt.Errorf("filtermap: next is not a chain")
	}

	choice := sfn.NewChoice(stack, jsii.String("choice"), &sfn.ChoiceProps{
		Comment:    jsii.String("choice"),
		OutputPath: jsii.String("$.__input"),
	})

	this := sfn.NewPass(stack, jsii.String("pass"), &sfn.PassProps{
		Parameters: &map[string]interface{}{
			"__input.$": "$",
		},
		Comment: jsii.String("save input to $.__input"),
	})

	function, err := s.Lambda.Compile(stack, choice, 0)
	if err != nil {
		return nil, fmt.Errorf("filter: %w", err)
	}

	this.Next(function[0].(sfn.IChainable))

	choice.When(sfn.Condition_BooleanEquals(jsii.String("$.__choice.result"), jsii.Bool(true)),
		next, &sfn.ChoiceTransitionOptions{},
	).Otherwise(
		sfn.NewSucceed(stack, jsii.String("filtered"), &sfn.SucceedProps{}),
	)

	return []interface{}{
		this,
	}, nil
}
