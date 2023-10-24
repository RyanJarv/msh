package event

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-sdk-go-v2/aws"
	"os"
)

var args = struct {
	Cron *string
	Rate *string
}{
	Cron: flag.String("cron", "", "cron expression (i.e: 0 18 ? * MON-FRI *)"),
	Rate: flag.String("rate", "", "rate expression (i.e: 1 minute, 5 hours, 1 day, etc...)"),
}

type Schedule string

func (s Schedule) ExpressionString() *string {
	return aws.String(string(s))
}

func NewEvent() (*Event, error) {
	props := &awsevents.RuleProps{
		EventPattern: &awsevents.EventPattern{},
	}

	if *args.Cron != "" {
		props.Schedule = Schedule(*args.Cron)
	} else if *args.Rate != "" {
		props.Schedule = awsevents.Schedule_Expression(aws.String(fmt.Sprintf("rate(%s)", *args.Rate)))
	} else if flag.NArg() < 2 {
		// Get the last argument from the end to support running with a shebang.
		path := flag.Arg(flag.NFlag())

		config, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("reading file: %s: %w", path, err)
		}

		if err := json.Unmarshal(config, &props.EventPattern); err != nil {
			return nil, fmt.Errorf("unmarshalling config: %w", err)
		}
	} else {
		return nil, fmt.Errorf("must provide either cron or rate expression or event pattern")
	}

	return &Event{
		RuleProp: props,
	}, nil
}

type Event struct {
	RuleProp *awsevents.RuleProps
}

func (s *Event) Name() string { return "event" }

func (s *Event) Run(stack awscdk.Stack, next interface{}) (interface{}, error) {
	target, ok := next.(awsevents.IRuleTarget)
	if !ok {
		return nil, fmt.Errorf("next step must be eventbridge target, got: %T", next)
	}

	rule := awsevents.NewRule(stack, aws.String("event"), s.RuleProp)
	rule.AddTarget(target)

	// Nothing can chain into an event rule.
	return nil, nil
}
