package sfn

import (
	_ "embed"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/jsii-runtime-go"
)

func NewSfn() (*Sfn, error) {
	return &Sfn{}, nil
}

type Sfn struct{}

func (s *Sfn) Name() string { return "statemachine" }

func (s *Sfn) Run(stack awscdk.Stack, last interface{}) (interface{}, error) {
	rule, ok := last.(awsevents.Rule)
	if !ok {
		return nil, fmt.Errorf("last step must be eventbridge rule, got: %T", last)
	}

	chain := sfn.Chain_Start(
		sfn.NewPass(stack, jsii.String("choice"), &sfn.PassProps{}),
	)

	machine := sfn.NewStateMachine(stack, jsii.String("StateMachine"), &sfn.StateMachineProps{
		DefinitionBody: sfn.DefinitionBody_FromChainable(chain),
		Timeout:        awscdk.Duration_Minutes(jsii.Number(5)),
		Comment:        jsii.String("a super cool state machine"),
	})

	rule.AddTarget(
		awseventstargets.NewSfnStateMachine(machine, &awseventstargets.SfnStateMachineProps{}),
	)

	return chain, nil
}
