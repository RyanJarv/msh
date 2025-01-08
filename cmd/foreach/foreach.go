package foreach

import (
	_ "embed"
	"fmt"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/types"
)

func New(_ *app.App) (*Foreach, error) {
	return &Foreach{}, nil
}

type Foreach struct {
	types.IIterator `json:"-"`
}

func (s Foreach) GetName() string { return "foreach" }

func (s *Foreach) Init(scope constructs.Construct, i int) error {
	name := fmt.Sprintf("%s-%d", s.GetName(), i)
	s.IIterator = sfn.NewMap(scope, jsii.String(name), &sfn.MapProps{
		Comment: jsii.String(name),
	})
	return nil
}
