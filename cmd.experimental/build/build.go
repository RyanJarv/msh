package build

import (
	_ "embed"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/ryanjarv/msh/pkg/app"
)

func New(_ *app.App) (*Build, error) {
	return &Build{}, nil
}

type Build struct {
	awsstepfunctions.IChainable
}

func (s Build) GetName() string { return "build" }

func (s *Build) Init(scope constructs.Construct, i int) error {
	s.IChainable = awsstepfunctions.NewSucceed(scope, jsii.String("build-succeed"), nil)

	return nil
}
