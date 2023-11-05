package cron

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
	"strings"
)

func New(app app.App) (*Cron, error) {
	opts := strings.Split(strings.Join(app.Args(), " "), " ")
	if len(opts) != 6 {
		return nil, fmt.Errorf("cron: must have 6 args, got: %d", len(flag.Args()))
	}

	cronopts := awsevents.CronOptions{
		Minute:  jsii.String(app.Arg(0)),
		Hour:    jsii.String(app.Arg(1)),
		Day:     jsii.String(app.Arg(2)),
		WeekDay: jsii.String(app.Arg(3)),
		Month:   jsii.String(app.Arg(4)),
		Year:    jsii.String(app.Arg(5)),
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
	awsevents.Rule
	CronOptions awsevents.CronOptions
}

func (s Cron) GetName() string { return "schedule" }

func (s Cron) Compile(stack constructs.Construct, i int) error {
	s.Rule = awsevents.NewRule(stack, aws.String(s.GetName()), &awsevents.RuleProps{
		Schedule: awsevents.Schedule_Cron(&s.CronOptions),
	})

	return nil
}
