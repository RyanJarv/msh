package sns

import (
	_ "embed"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/ryanjarv/msh/pkg/app"
)

func New(_ app.App) (*Build, error) {
	return &Build{}, nil
}

type Build struct{}

func (s Build) GetName() string { return "build" }

func (s Build) Compile(_ constructs.Construct, _ []interface{}, i int) ([]interface{}, error) {
	return []interface{}{}, nil
}
