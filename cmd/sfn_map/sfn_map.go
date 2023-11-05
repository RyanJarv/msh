package sfn_map

import (
	_ "embed"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/types"
)

func New(app app.App) (*Map, error) {
	return &Map{}, nil
}

type Map struct {
	types.IChain
	types.IIterator
	sfn.Map
}

func (s Map) GetName() string { return "map" }

func (s *Map) Compile(stack constructs.Construct, i int) error {
	s.Map = sfn.NewMap(stack, jsii.String("map"), &sfn.MapProps{
		Comment: jsii.String("map"),
	})

	return nil
}
