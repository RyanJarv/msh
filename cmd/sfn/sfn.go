package sfn

import (
	_ "embed"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/types"
	"strings"
)

func New(app *app.App) (*Sfn, error) {
	name := app.StackName()
	name = strings.Replace(name, "-", "_", -1)

	return &Sfn{
		Name: fmt.Sprintf("msh_sfn_%s", name),
	}, nil
}

type Sfn struct {
	Name         string
	StateMachine sfn.IStateMachine
}

func (s *Sfn) GetName() string { return "sfn" }

func (s *Sfn) Run(step *types.StepRunInfo) error {
	s.StateMachine = sfn.NewStateMachine(step.Scope, step.Id("sfn"), &sfn.StateMachineProps{
		StateMachineName: jsii.String(s.Name),
		DefinitionBody:   sfn.NewChainDefinitionBody(sfn.Chain_Start(s.GetChain(step))),
		Timeout:          awscdk.Duration_Minutes(jsii.Number(5)),
		Comment:          jsii.String("a super cool app machine"),
	})

	return nil
}

// GetChain returns a chain defining the full sfn definition. The Chain can't be modified after the
// state machine is created.
func (s *Sfn) GetChain(step *types.StepRunInfo) sfn.IChainable {
	if step == nil {
		panic(fmt.Errorf("step is nil"))
	} else if step.Next == nil {
		return sfn.NewSucceed(step.Scope, step.Id("success"), &sfn.SucceedProps{})
	} else if v, ok := step.Next.Step.(types.SfnChainable); ok {
		return sfn.Chain_Start(v.GetChain(step.Next))
	}

	return nil
}

// Bind fulfills the awsevents.IRuleTarget interface.
// Note: Index'm not sure why but cdk doesn't let us just embed this in the Sfn struct.
func (s *Sfn) Bind(rule awsevents.IRule, id *string) *awsevents.RuleTargetConfig {
	return awseventstargets.NewSfnStateMachine(s.StateMachine, &awseventstargets.SfnStateMachineProps{}).Bind(rule, id)
}
