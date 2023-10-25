package each

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/types"
	"os"
)

func New(registry types.Registry) (*Each, error) {
	apps := []app.App{}

	for _, path := range flag.Args() {
		f, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("opening file: %s: %w", path, err)
		}

		app, err := app.GetPipeline(registry, f, nil)
		if err != nil {
			return nil, fmt.Errorf("get app: %w", err)
		}
		apps = append(apps, app)

	}
	return &Each{
		SubApp: apps,
	}, nil
}

type Each struct {
	SubApp []app.App
}

func (s Each) GetName() string { return "each" }

func (s Each) Compile(stack awscdk.Stack, next []interface{}) ([]interface{}, error) {
	var heads []interface{}
	for _, subapp := range s.SubApp {
		chain, err := subapp.Compile(stack, next)
		if err != nil {
			return nil, fmt.Errorf("each: %w", err)
		}

		heads = append(heads, chain[0])
	}

	return heads, nil
}
