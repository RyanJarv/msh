package sfn_map

import (
	_ "embed"
	"fmt"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
)

func New(app app.App) (*Map, error) {
	return &Map{}, nil
}

type Map struct{}

func (s Map) GetName() string { return "map" }

func (s Map) Compile(stack constructs.Construct, _next interface{}, i int) ([]interface{}, error) {
	next, ok := _next.(sfn.IChainable)
	if !ok {
		return nil, fmt.Errorf("next is not a chain: got %T", _next)
	}

	sfn_map := sfn.NewMap(stack, jsii.String("map"), &sfn.MapProps{
		Comment: jsii.String("map"),
	})
	sfn_map.Iterator(next)

	return []interface{}{
		sfn_map,
	}, nil
}
