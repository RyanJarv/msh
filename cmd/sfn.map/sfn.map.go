package sfn_map

import (
	_ "embed"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/types"
)

func New(registry types.Registry) (*Map, error) {
	return &Map{}, nil
}

type Map struct{}

func (s Map) GetName() string { return "map" }

func (s Map) Compile(stack awscdk.Stack, _next []interface{}) ([]interface{}, error) {
	next, ok := _next[0].(sfn.IChainable)
	if !ok {
		return nil, fmt.Errorf("filtermap: next is not a chain")
	}

	sfn_map := sfn.NewMap(stack, jsii.String("map"), &sfn.MapProps{
		Comment: jsii.String("map"),
	})
	sfn_map.Iterator(next)

	return []interface{}{
		sfn_map,
	}, nil
}
