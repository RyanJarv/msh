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

func (s *Sfn) Name() string { return "statemachine" }

func (s *Sfn) Run(stack awscdk.Stack, last interface{}) (interface{}, error) {
	var ok bool
	s.rule, ok = last.(awsevents.Rule)
	if !ok {
		return nil, fmt.Errorf("last step must be eventbridge rule, got: %T", last)
	}

	s.Chain = sfn.NewPass(stack, jsii.String("pass"), &sfn.PassProps{})

	return s.Chain, nil
}

func (s *Sfn) Finalize(stack awscdk.Stack) error {
	// The state machine must be created after the chain is set up otherwise we won't see all the steps.
	machine := sfn.NewStateMachine(stack, jsii.String("StateMachine"), &sfn.StateMachineProps{
		DefinitionBody: sfn.DefinitionBody_FromChainable(s.Chain),
		Timeout:        awscdk.Duration_Minutes(jsii.Number(5)),
		Comment:        jsii.String("a super cool state machine"),
	})

	target := awseventstargets.NewSfnStateMachine(machine, &awseventstargets.SfnStateMachineProps{})
	s.rule.AddTarget(target)

	return nil
}
