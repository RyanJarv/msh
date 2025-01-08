package call

import (
	_ "embed"
	"fmt"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	tasks "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
)

func New(app *app.App) (*Call, error) {
	return &Call{
		Name: app.Flag.Arg(1),
	}, nil
}

type Call struct {
	Name string
	tasks.StepFunctionsStartExecution
}

func (s Call) GetName() string { return "ref" }

func (s *Call) Init(scope constructs.Construct, i int) error {
	name := fmt.Sprintf("%s-%s-%d", s.GetName(), s.Name, i)

	props := tasks.StepFunctionsStartExecutionProps{
		StateMachine: sfn.StateMachine_FromStateMachineName(scope, jsii.String(name), jsii.String(s.Name)),
	}

	s.StepFunctionsStartExecution = tasks.NewStepFunctionsStartExecution(scope, jsii.String(fmt.Sprintf("invoke-%s", name)), &props)

	return nil
}
