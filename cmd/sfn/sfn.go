package sfn

import (
	_ "embed"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/utils"
)

func New(app app.App) (*Sfn, error) {
	flags := utils.ParseArgs(app.Args)

	name := flags.Arg(1)
	if name == "" {
		name = "sfn"
	}

	return &Sfn{
		Name: name,
	}, nil
}

type Sfn struct {
	Name string
	awsevents.IRuleTarget
}

func (s Sfn) GetName() string { return "sfn" }

func (s *Sfn) Compile(stack constructs.Construct, next interface{}, i int) ([]interface{}, error) {
	chain, ok := next.(sfn.IChainable)
	if !ok {
		return nil, fmt.Errorf("next step must be statemachine task, got: %T", next)
	}

	// The app machine must be created after the chain is set up otherwise we won't see all the steps.
	machine := sfn.NewStateMachine(stack, jsii.String(s.Name), &sfn.StateMachineProps{
		StateMachineName: jsii.String(s.Name),
		DefinitionBody:   sfn.DefinitionBody_FromChainable(chain),
		Timeout:          awscdk.Duration_Minutes(jsii.Number(5)),
		Comment:          jsii.String("a super cool app machine"),
	})

	target := awseventstargets.NewSfnStateMachine(machine, &awseventstargets.SfnStateMachineProps{})
	return []interface{}{target}, nil
}
