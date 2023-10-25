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

type Sfn struct {
	rule  awsevents.Rule
	Chain sfn.IChainable
}

func (s Sfn) GetName() string { return "statemachine" }

func (s Sfn) Compile(stack awscdk.Stack, next interface{}) (interface{}, error) {
	chain, ok := next.(sfn.IChainable)
	if !ok {
		return nil, fmt.Errorf("next step must be statemachine task, got: %T", next)
	}

	pass := sfn.NewPass(stack, jsii.String("pass"), &sfn.PassProps{})
	pass.Next(chain)

	// The app machine must be created after the chain is set up otherwise we won't see all the steps.
	machine := sfn.NewStateMachine(stack, jsii.String("StateMachine"), &sfn.StateMachineProps{
		DefinitionBody: sfn.DefinitionBody_FromChainable(s.Chain),
		Timeout:        awscdk.Duration_Minutes(jsii.Number(5)),
		Comment:        jsii.String("a super cool app machine"),
	})

	return awseventstargets.NewSfnStateMachine(machine, &awseventstargets.SfnStateMachineProps{}), nil
}
