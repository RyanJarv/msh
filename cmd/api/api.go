package api

import (
	_ "embed"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	tasks "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/jsii-runtime-go"
)

func New() (*Api, error) {
	return &Api{}, nil
}

type Api struct {
	rule awsevents.Rule
}

func (s Api) GetName() string { return "sfn" }

func (s Api) Compile(stack awscdk.Stack, next interface{}) ([]interface{}, error) {
	var chain sfn.IChainable
	if next != nil {
		var ok bool
		chain, ok = next.(sfn.IChainable)
		if !ok {
			return nil, fmt.Errorf("next step must be chainable")
		}
	} else {
		chain = sfn.NewSucceed(stack, jsii.String("succeed"), &sfn.SucceedProps{})
	}

	noop := sfn.NewPass(stack, jsii.String("pass"), &sfn.PassProps{})

	iter := sfn.NewMap(stack, jsii.String("map"), &sfn.MapProps{
		InputPath:  sfn.JsonPath_StringAt(jsii.String("$.Reservations[*].Instances[*]")),
		ResultPath: sfn.JsonPath_DISCARD(),
	})

	call := tasks.NewCallAwsService(stack, jsii.String("call"), &tasks.CallAwsServiceProps{
		Service:      jsii.String("ec2"),
		Action:       jsii.String("describeInstances"),
		IamResources: jsii.Strings("*"),
		Parameters: &map[string]interface{}{
			"MaxResults": jsii.Number(5),
		},
	}).Next(iter)

	more := tasks.NewCallAwsService(stack, jsii.String("more"), &tasks.CallAwsServiceProps{
		Service:      jsii.String("ec2"),
		Action:       jsii.String("describeInstances"),
		IamResources: jsii.Strings("*"),
		Parameters: &map[string]interface{}{
			"NextToken.$": jsii.String("$.NextToken"),
			"MaxResults":  jsii.Number(5),
		},
	}).Next(iter)

	iter.Iterator(noop).
		Next(sfn.NewChoice(stack, jsii.String("choice"), &sfn.ChoiceProps{}).
			When(sfn.Condition_IsPresent(jsii.String("$.NextToken")), more, &sfn.ChoiceTransitionOptions{}).
			Otherwise(chain))

	return []interface{}{call}, nil
}
