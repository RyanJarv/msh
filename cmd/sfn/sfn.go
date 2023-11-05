package sfn

import (
	_ "embed"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	tasks "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
)

func New(app app.App) (*Sfn, error) {
	name := app.Arg(1)
	if name == "" {
		name = "sfn"
	}

	return &Sfn{
		Name: name,
	}, nil
}

type Sfn struct {
	awsevents.IRuleTarget
	sfn.INextable
	Name    string
	Chain   sfn.IChainable
	machine sfn.StateMachine
}

func (s Sfn) GetName() string { return "sfn" }

func (s *Sfn) Compile(stack constructs.Construct, i int) error {
	// The app machine must be created after the chain is set up otherwise we won't see all the steps.
	s.machine = sfn.NewStateMachine(stack, jsii.String(s.Name), &sfn.StateMachineProps{
		StateMachineName: jsii.String(s.Name),
		DefinitionBody:   sfn.NewChainDefinitionBody(sfn.NewPass(stack, jsii.String("this-should-be-overridden"), &sfn.PassProps{})),
		Timeout:          awscdk.Duration_Minutes(jsii.Number(5)),
		Comment:          jsii.String("a super cool app machine"),
	})

	s.IRuleTarget = awseventstargets.NewSfnStateMachine(s.machine, &awseventstargets.SfnStateMachineProps{})

	return nil
}

func (s *Sfn) Next(chain sfn.IChainable) sfn.Chain {
	sfn.NewStateMachine_Override(s.machine, s.machine.Stack(), jsii.String(s.Name), &sfn.StateMachineProps{
		DefinitionBody: sfn.DefinitionBody_FromChainable(chain),
	})

	props := tasks.StepFunctionsStartExecutionProps{
		StateMachine: s.machine,
	}

	invoke := tasks.NewStepFunctionsStartExecution(s.machine.Stack(), jsii.String(fmt.Sprintf("invoke-%s", s.Name)), &props)
	return invoke.Next(chain)
}
