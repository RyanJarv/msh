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
	"syscall"
)

func New(_ app.App) (*Sfn, error) {
	return &Sfn{}, nil
}

type Sfn struct {
	stack   constructs.Construct
	machine sfn.StateMachine
	name    string
}

func (s *Sfn) GetName() string { return "sfn" }

func (s *Sfn) Compile(stack constructs.Construct, i int) error {
	s.stack = stack
	s.name = fmt.Sprintf("%s-%d", s.GetName(), syscall.Getpid())
	return nil
}

func (s *Sfn) Next(chain sfn.IChainable) sfn.Chain {
	// The app machine must be created after the chain is set up otherwise we won't see all the steps.
	s.machine = sfn.NewStateMachine(s.stack, jsii.String(s.name), &sfn.StateMachineProps{
		StateMachineName: jsii.String(s.name),
		DefinitionBody:   sfn.NewChainDefinitionBody(chain),
		Timeout:          awscdk.Duration_Minutes(jsii.Number(5)),
		Comment:          jsii.String("a super cool app machine"),
	})

	return nil
}

// Bind fulfills the awsevents.IRuleTarget interface.
// Note: I'm not sure why but cdk doesn't let us just embed this in the Sfn struct.
func (s *Sfn) Bind(rule awsevents.IRule, id *string) *awsevents.RuleTargetConfig {
	return awseventstargets.NewSfnStateMachine(s.machine, &awseventstargets.SfnStateMachineProps{}).Bind(rule, id)
}
