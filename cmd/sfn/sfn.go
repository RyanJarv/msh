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

func New(app *app.App, _ []string) (types.CdkStep, error) {
	name := strings.Replace(app.StackName(), "-", "_", -1)

	return &Sfn{
		Name: fmt.Sprintf("msh_sfn_%s", name),
	}, nil
}

type Sfn struct {
	Name         string
	StateMachine sfn.IStateMachine
}

func (s *Sfn) GetName() string { return "sfn" }

func (s *Sfn) Run(i *types.StepRunInfo) error {
	s.StateMachine = sfn.NewStateMachine(i.Scope, i.Id("sfn"), &sfn.StateMachineProps{
		StateMachineName: jsii.String(s.Name),
		DefinitionBody:   sfn.NewChainDefinitionBody(sfn.Chain_Start(i.GetChain())),
		Timeout:          awscdk.Duration_Minutes(jsii.Number(5)),
		Comment:          jsii.String("a super cool app machine"),
	})

	return nil
}

// Bind fulfills the awsevents.IRuleTarget interface.
// Note: Index'm not sure why but cdk doesn't let us just embed this in the Sfn struct.
func (s *Sfn) Bind(rule awsevents.IRule, id *string) *awsevents.RuleTargetConfig {
	return awseventstargets.NewSfnStateMachine(s.StateMachine, &awseventstargets.SfnStateMachineProps{}).Bind(rule, id)
}
