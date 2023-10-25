package each

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/ryanjarv/msh/pkg/app"
	"os"
)

func New(a app.App) (*Each, error) {
	each := &Each{
		SubApp: []app.App{},
	}

	for _, path := range flag.Args()[1:] {
		f, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("opening file: %s: %w", path, err)
		}

		subapp, err := app.GetApp(a.Registry, f, nil)
		if err != nil {
			return nil, fmt.Errorf("getting subapp: %w", err)
		}

		each.SubApp = append(each.SubApp, subapp)
	}

	return each, nil
}

type Each struct {
	SubApp []app.App
}

func (s *Each) Name() string { return "each" }

func (s *Each) Compile(stack awscdk.Stack, next []interface{}) ([]interface{}, error) {
	var heads []interface{}
	for _, s := range s.SubApp {
		chain, err := s.Compile(stack, next)
		if err != nil {
			return nil, fmt.Errorf("each: %w", err)
		}

		heads = append(heads, chain[0])
	}

	return heads, nil
}
