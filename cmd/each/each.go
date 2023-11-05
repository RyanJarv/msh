package each

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/ryanjarv/msh/pkg/app"
	"os"
)

func New(a app.App) (*Each, error) {
	apps := []app.App{}

	for _, path := range flag.Args() {
		f, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("opening file: %s: %w", path, err)
		}

		app, err := app.GetPipeline(a.Registry, f, nil, nil)
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

func (s Each) Compile(stack constructs.Construct, next []interface{}, i int) ([]interface{}, error) {
	var heads []interface{}
	for _, subapp := range s.SubApp {
		chain, err := subapp.Compile(stack, next, 0)
		if err != nil {
			return nil, fmt.Errorf("each: %w", err)
		}

		heads = append(heads, chain[0])
	}

	return heads, nil
}
