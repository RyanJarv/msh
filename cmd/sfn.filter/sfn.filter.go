package sfn_filter

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/types"
	"os"
)

func New(registry types.Registry) (*Filter, error) {
	if flag.NFlag() > 1 {
		return nil, fmt.Errorf("if: only one flag allowed")
	}

	path := flag.Arg(0)
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening file: %s: %w", path, err)
	}

	filterMap := Filter{}

	cond, err := app.GetPipeline(registry, f, nil)
	if err != nil {
		return nil, fmt.Errorf("get cond: %w", err)
	}

	filterMap.SubApp = &cond

	return &filterMap, nil
}

type Filter struct {
	SubApp *app.App
}

func (s Filter) GetName() string { return "if" }

func (s Filter) Compile(stack awscdk.Stack, _next []interface{}) ([]interface{}, error) {
	next, ok := _next[0].(sfn.IChainable)
	if !ok {
		return nil, fmt.Errorf("filtermap: next is not a chain")
	}

	choice := sfn.NewChoice(stack, jsii.String("choice"), &sfn.ChoiceProps{
		Comment: jsii.String("choice"),
	})
	_, err := s.SubApp.Compile(stack, []interface{}{choice})
	if err != nil {
		return nil, fmt.Errorf("if: %w", err)
	}

	choice.When(sfn.Condition_BooleanEquals(jsii.String("$.result"), jsii.Bool(true)),
		next, &sfn.ChoiceTransitionOptions{},
	).Otherwise(
		sfn.NewSucceed(stack, jsii.String("filtered"), &sfn.SucceedProps{}),
	)

	return []interface{}{
		s.SubApp.State.Steps[0].Value,
	}, nil
}
