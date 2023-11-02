package call

import (
	_ "embed"
	"flag"
	"fmt"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	tasks "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func New(args []string) (*Call, error) {
	flagset := flag.NewFlagSet("sfn", flag.ExitOnError)
	err := flagset.Parse(args)
	if err != nil {
		return nil, fmt.Errorf("sfn: %w", err)
	}

	return &Call{
		Name: flagset.Arg(1),
	}, nil
}

type Call struct {
	Name string
}

func (s Call) GetName() string { return "ref" }

func (s Call) Compile(stack constructs.Construct, next interface{}, i int) ([]interface{}, error) {
	invoke := tasks.NewStepFunctionsStartExecution(stack, jsii.String(fmt.Sprintf("invoke-%s-%s", s.GetName(), s.Name)), &tasks.StepFunctionsStartExecutionProps{
		StateMachine: sfn.StateMachine_FromStateMachineName(stack, jsii.String(fmt.Sprintf("%s-%s", s.GetName(), s.Name)), jsii.String(s.Name)),
	})

	if chain, ok := next.(sfn.IChainable); ok {
		invoke.Next(chain)
	}

	return []interface{}{invoke}, nil
}
