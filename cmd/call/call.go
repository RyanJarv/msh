package call

import (
	_ "embed"
	"fmt"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	tasks "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/utils"
)

func New(app app.App) (*Call, error) {
	flags := utils.ParseArgs(app.Args)

	return &Call{
		Name: flags.Arg(1),
	}, nil
}

type Call struct {
	Name string
}

func (s Call) GetName() string { return "ref" }

func (s Call) Compile(stack constructs.Construct, next interface{}, i int) ([]interface{}, error) {
	name := fmt.Sprintf("%s-%s-%d", s.GetName(), s.Name, i)

	props := tasks.StepFunctionsStartExecutionProps{
		StateMachine: sfn.StateMachine_FromStateMachineName(stack, jsii.String(name), jsii.String(s.Name)),
	}

	invoke := tasks.NewStepFunctionsStartExecution(stack, jsii.String(fmt.Sprintf("invoke-%s", name)), &props)

	if chain, ok := next.(sfn.IChainable); ok {
		invoke.Next(chain)
	}

	return []interface{}{invoke}, nil
}
