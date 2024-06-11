package exclusive

import (
	_ "embed"
	"fmt"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	tasks "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/types"
)

func New(app app.App) (*Call, error) {
	return &Call{
		Name: app.Arg(1),
	}, nil
}

type Call struct {
	Name string
	tasks.CallAwsService
	types.IChain
}

func (s Call) GetName() string { return "sfn" }

func (s *Call) Compile(stack constructs.Construct, i int) error {
	name := fmt.Sprintf("%s-%s-%d", s.GetName(), s.Name, i)

	s.IChain = ListStateMachines(stack, name, false).
		Next(
			sfn.NewMap(stack, jsii.String("map-sfn"),
				&sfn.MapProps{
					ItemsPath: sfn.JsonPath_StringAt(jsii.String("$.Executions")),
				},
			).Iterator(
				sfn.NewChoice(
					stack, jsii.String("choice-sfn"),
					&sfn.ChoiceProps{
						Comment: jsii.String("Check if the execution is the current one"),
					},
				).When(
					sfn.Condition_StringEqualsJsonPath(jsii.String("$.ExecutionArn"), jsii.String("$$.Execution.Id")),
					sfn.NewPass(stack, jsii.String("pass-sfn"), &sfn.PassProps{}),
					&sfn.ChoiceTransitionOptions{},
				),
			),
		).
		Next(
			sfn.NewChoice(
				stack, jsii.String("choice-sfn-paginator"),
				&sfn.ChoiceProps{
					Comment: jsii.String("Check if there are more pages to fetch"),
				},
			).When(
				sfn.Condition_IsPresent(jsii.String("$.NextToken")),
				ListStateMachines(stack, "sfn-paginator", true),
				&sfn.ChoiceTransitionOptions{},
			).Otherwise(
				sfn.NewPass(stack, jsii.String("sfn-done"), &sfn.PassProps{}),
			),
		)

	return nil
}

func ListStateMachines(stack constructs.Construct, name string, next bool) tasks.CallAwsService {
	params := map[string]interface{}{
		"StateMachineArn.$": sfn.JsonPath_StringAt(jsii.String("$$.StateMachine.Id")),
		"StatusFilter":      jsii.String("RUNNING"),
	}
	if next {
		params["NextToken.$"] = sfn.JsonPath_StringAt(jsii.String("$.NextToken"))
	}

	return tasks.NewCallAwsService(stack, jsii.String(fmt.Sprintf("invoke-%s", name)),
		&tasks.CallAwsServiceProps{
			Service:    jsii.String("states"),
			Action:     jsii.String("ListStateMachines"),
			Parameters: &params,
		})
}
