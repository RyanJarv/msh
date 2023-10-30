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
	lambda, err := lambda.New(args)
	if err != nil {
		return nil, err
	}

	return &Filter{
		Lambda: lambda,
	}, nil
}

type Filter struct {
	Lambda *lambda.LambdaCmd
}

func (s Filter) GetName() string { return "filter" }

func (s Filter) Compile(stack constructs.Construct, _next interface{}) ([]interface{}, error) {
	next, ok := _next.(sfn.IChainable)
	if !ok {
		return nil, fmt.Errorf("filtermap: next is not a chain")
	}

	choice := sfn.NewChoice(stack, jsii.String("choice"), &sfn.ChoiceProps{
		Comment: jsii.String("choice"),
	})

	this, err := s.Lambda.Compile(stack, choice)
	if err != nil {
		return nil, fmt.Errorf("filter: %w", err)
	}

	choice.When(sfn.Condition_BooleanEquals(jsii.String("$.result"), jsii.Bool(true)),
		next, &sfn.ChoiceTransitionOptions{},
	).Otherwise(
		sfn.NewSucceed(stack, jsii.String("filtered"), &sfn.SucceedProps{}),
	)

	return this, nil
}
