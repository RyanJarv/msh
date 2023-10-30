package event

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/ryanjarv/msh/pkg/utils"
	"github.com/samber/lo"
	"os"
	"strings"
)

func NewEvent(args []string) (*Event, error) {
	//args := flag.Args()

	//args, err := ParseTemplateArgs(args)
	//if err != nil {
	//	return nil, fmt.Errorf("parse template args: %w", err)
	//}

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

		//strings.TrimPrefix()
	}

	return nil, nil, fmt.Errorf("parse template args: failed to find end of template args")
}

type Event struct {
	RuleProp *awsevents.RuleProps
}

func (s Event) GetName() string { return "event" }

func (s Event) Compile(stack constructs.Construct, next []interface{}) ([]interface{}, error) {
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
