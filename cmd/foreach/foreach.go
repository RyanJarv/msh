package foreach

import (
	_ "embed"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/types"
)

func New(app app.App) (*Foreach, error) {
	return &Foreach{}, nil
}

type Foreach struct {
	types.IChain
	types.IIterator
	sfn.Map
}

func (s Foreach) GetName() string { return "map" }

func (s *Foreach) Compile(stack constructs.Construct, i int) error {
	s.Map = sfn.NewMap(stack, jsii.String("map"), &sfn.MapProps{
		Comment: jsii.String("map"),
	})

	return nil
}
