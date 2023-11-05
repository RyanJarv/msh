package event

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/samber/lo"
	"os"
	"strings"
)

func New(_ app.App) (*Event, error) {
	path := flag.Arg(flag.NFlag())
	config, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %s: %w", path, err)
	}

	event := &awsevents.EventPattern{}

	if err := json.Unmarshal(config, &event); err != nil {
		return nil, fmt.Errorf("unmarshalling config: %w", err)
	}

	return &Event{
		RuleProp: &awsevents.RuleProps{
			EventPattern: event,
		},
	}, nil
}

func ParseTemplateArgs(args []string) (map[string]any, []string, error) {
	_ = map[string]any{}

	start := lo.IndexOf(args, "flag{")
	if start == -1 {
		return nil, args, nil
	}

	end := lo.IndexOf(args[start:], "}")
	if start == -1 {
		return nil, nil, fmt.Errorf("expected `}` got: %s", args[start:])
	}

	for _, arg := range args[start:end] {
		parts := strings.Split(arg, "=")
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("template flags must have a single `=` got: %s", arg)
		}
	}

	return nil, nil, fmt.Errorf("parse template args: failed to find end of template args")
}

type Event struct {
	awsevents.Rule
	RuleProp *awsevents.RuleProps
}

func (s Event) GetName() string { return "event" }

func (s Event) Compile(stack constructs.Construct, i int) error {
	s.Rule = awsevents.NewRule(stack, aws.String("event"), s.RuleProp)

	return nil
}
