package foreach

import (
	_ "embed"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/types"
)

func New(_ *app.App, _ []string) (types.CdkStep, error) {
	return &Foreach{}, nil
}

type Foreach struct{}

func (s Foreach) GetName() string { return "foreach" }

func (s *Foreach) GetChain(i *types.StepRunInfo) sfn.IChainable {
	return sfn.NewMap(i.Scope, i.Id("map"), &sfn.MapProps{
		Comment: i.Id("map"),
	}).
		Iterator(i.GetChain()).
		StartState()
}
