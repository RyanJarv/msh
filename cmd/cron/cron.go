package cron

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/logs"
	"github.com/ryanjarv/msh/pkg/types"
	"github.com/samber/lo"
	"strings"
)

func New(app *app.App, argv []string) (types.CdkStep, error) {
	opts := flag.NewFlagSet("cron", flag.ExitOnError)
	opts.Parse(argv[1:])

	args := strings.Split(strings.Join(opts.Args(), " "), " ")
	args = lo.Reject(args, func(item string, _ int) bool {
		return item == ""
	})
	if len(args) != 6 {
		return nil, fmt.Errorf("cron: must have 6 args, got: %d", len(args))
	}

	cronopts := awsevents.CronOptions{
		Minute:  jsii.String(args[0]),
		Hour:    jsii.String(args[1]),
		Day:     jsii.String(args[2]),
		WeekDay: jsii.String(args[3]),
		Month:   jsii.String(args[4]),
		Year:    jsii.String(args[5]),
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
		return fmt.Errorf("run: next must be IRuleTarget, got: %T %+v", step.Next, step.Next)
	}

	return nil
}
