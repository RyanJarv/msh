package cron

import (
	_ "embed"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/logs"
	"github.com/ryanjarv/msh/pkg/types"
	"strings"
)

func New(app *app.App) (*Cron, error) {
	opts := strings.Split(strings.Join(app.Flag.Args()[1:], " "), " ")
	if len(opts) != 6 {
		return nil, fmt.Errorf("cron: must have 6 args, got: %d", app.Flag.NArg())
	}

	cronopts := awsevents.CronOptions{
		Minute:  jsii.String(opts[0]),
		Hour:    jsii.String(opts[1]),
		Day:     jsii.String(opts[2]),
		WeekDay: jsii.String(opts[3]),
		Month:   jsii.String(opts[4]),
		Year:    jsii.String(opts[5]),
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
	id          string
}

func (s *Cron) GetName() string { return "schedule" }

func (s *Cron) AfterRun(step *types.StepRunInfo) error {
	id := fmt.Sprintf("%s-%d", s.GetName(), step.Index)
	logs.Debug("Cron.Run: got scope id %s", *step.Scope.ToString())

	switch v := step.Next.Step.(type) {
	case awsevents.IRuleTarget:
		awsevents.NewRule(step.Scope, jsii.String("rule-"+id), &awsevents.RuleProps{
			Schedule: awsevents.Schedule_Cron(&s.CronOptions),
		}).AddTarget(v)
	default:
		return fmt.Errorf("run: next must be IRuleTarget or IStateMachine, got: %T %+v", step.Next, step.Next)
	}

	return nil
}
