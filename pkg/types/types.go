package types

import (
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

type Step struct {
	Name  string
	Value interface{}
}

func (s Step) GetName() string {
	return s.Name
}

type CdkStep interface {
	GetName() string
}

type BeforeRunCdkStep interface {
	BeforeRun(*StepRunInfo) error
}

type RunnableCdkStep interface {
	Run(*StepRunInfo) error
}

type AfterRunCdkStep interface {
	AfterRun(*StepRunInfo) error
}

type AcceptsLastRole interface {
	LastFoundRole(scope constructs.Construct, role awsiam.IRole)
}

type CdkStepCanInit interface {
	Init(scope constructs.Construct, i int) error
}

type SfnChainable interface {
	GetChain(*StepRunInfo) sfn.IChainable
}

type IChain interface {
	sfn.INextable
	sfn.IChainable
}

type IIterator interface {
	IChain
	Iterator(iterator sfn.IChainable) sfn.Map
}

type HasRole interface {
	Role() awsiam.IRole
}

type RequiresPermissions interface {
	RequiredPermissions(stack constructs.Construct) awsiam.PolicyStatement
}

type StepRunInfo struct {
	Index int
	Scope constructs.Construct
	Step  CdkStep
	Next  *StepRunInfo
	Last  *StepRunInfo
	Id    func(s ...string) *string
}

func (i StepRunInfo) GetChain() sfn.IChainable {
	if i.Next == nil || i.Next.Step == nil {
		return sfn.NewSucceed(i.Scope, i.Id("success"), &sfn.SucceedProps{})
	}
	switch v := i.Next.Step.(type) {
	case SfnChainable:
		return v.GetChain(i.Next).StartState()
	default:
		panic(fmt.Errorf("%s: next step must be SfnChainable, got: %T", i.Step.GetName(), i.Next.Step))
	}
}
