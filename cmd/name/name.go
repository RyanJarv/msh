package name

import (
	_ "embed"
	"github.com/ryanjarv/msh/pkg/app"
	L "github.com/ryanjarv/msh/pkg/logger"
)

func New(app *app.App) (*Name, error) {
	name := app.Flag.Arg(1)
	if name == "" {
		L.Error.Fatalf("Usage: %s <name>", app.Flag.Arg(1))
	}

	L.Debug.Printf("set CDK stack name: %s", name)
	app.State.Name = name
	return &Name{}, nil
}

type Name struct{}

func (s *Name) GetName() string { return "name" }
