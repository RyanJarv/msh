package cron

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/utils"
	"strings"
)

func New() (*Cron, error) {
	args := strings.Split(strings.Join(flag.Args(), " "), " ")
	if len(args) != 6 {
		return nil, fmt.Errorf("cron: must have 6 args, got: %d", len(args))
	}

	cronopts := awsevents.CronOptions{
		Minute:  &args[0],
		Hour:    &args[1],
		Day:     &args[2],
		WeekDay: &args[3],
		Month:   &args[4],
		Year:    &args[5],
	}

	// Cdk doesn't like when both day and weekday are specified.
	// If the value is "*", it's the same as not specifying it at all.
	if *cronopts.Day == "*" || *cronopts.Day == "?" {
		cronopts.Day = nil
	}

	if *cronopts.WeekDay == "*" || *cronopts.WeekDay == "?" {
		cronopts.WeekDay = nil
	}

	return &Cron{
		CronOptions: cronopts,
	}, nil
}

type Cron struct {
	CronOptions awsevents.CronOptions
}

func (s Cron) GetName() string { return "schedule" }

func (s Cron) Compile(stack constructs.Construct, next []interface{}, i int) ([]interface{}, error) {
	target, ok := utils.EachAs[awsevents.IRuleTarget](next)
	if !ok {
		return nil, fmt.Errorf("next step must be eventbridge target, got: %T: %+v", next[0], next[0])
	}

	rule := awsevents.NewRule(stack, aws.String(s.GetName()), &awsevents.RuleProps{
		Schedule: awsevents.Schedule_Cron(&s.CronOptions),
	})

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
	return jsii.String(fmt.Sprintf("cron(%s)", string(s)))
}
