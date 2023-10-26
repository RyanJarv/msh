package event

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/ryanjarv/msh/pkg/utils"
	"strings"
)

func NewCron() (*Cron, error) {
	return &Cron{
		RuleProp: &awsevents.RuleProps{
			EventPattern: &awsevents.EventPattern{},
			Schedule:     NewCronExpression(flag.Args()[1:]...),
		},
	}, nil
}

type Cron struct {
	RuleProp *awsevents.RuleProps
}

func (s Cron) GetName() string { return "event" }

func (s Cron) Compile(stack awscdk.Stack, next []interface{}) ([]interface{}, error) {
	target, ok := utils.EachAs[awsevents.IRuleTarget](next)
	if !ok {
		return nil, fmt.Errorf("next step must be eventbridge target, got: %T: %+v", next[0], next[0])
	}

	rule := awsevents.NewRule(stack, aws.String("event"), s.RuleProp)

	for _, t := range target {
		rule.AddTarget(t)
	}

	// Nothing can chain into an event rule.
	return nil, nil
}

func NewCronExpression(args ...string) CronExpression {
	return CronExpression(strings.Join(args, " "))
}

type CronExpression string

func (s CronExpression) ExpressionString() *string {
	return aws.String(fmt.Sprintf("cron(%s)", string(s)))
}
