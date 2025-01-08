package types

import (
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"reflect"
)

func NewRegistry(steps ...interface{}) Registry {
	r := Registry{}
	for _, step := range steps {
		f := reflect.TypeOf(step)
		if f.Kind() != reflect.Func {
			panic("step must be func")
		}

		v := f.Out(0)
		out := reflect.New(v.Elem())
		value := out.Interface()

		if v, ok := value.(CdkStep); !ok {
			panic(fmt.Sprintf("step must return CdkStep, got: %T %+v", value, value))
		} else {
			r[v.GetName()] = v
		}
	}

	return r
}

type Registry map[string]CdkStep

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
	BeforeRun(last *StepRunInfo) error
}

type RunnableCdkStep interface {
	Run(last *StepRunInfo) error
}

type AfterRunCdkStep interface {
	AfterRun(last *StepRunInfo) error
}

type AcceptsLastRole interface {
	LastFoundRole(scope constructs.Construct, role awsiam.IRole)
}

type CdkStepCanInit interface {
	Init(scope constructs.Construct, i int) error
}

type SfnChainable interface {
	GetChain(*StepRunInfo) awsstepfunctions.IChainable
}

type IChain interface {
	awsstepfunctions.INextable
	awsstepfunctions.IChainable
}

type IIterator interface {
	IChain
	Iterator(iterator awsstepfunctions.IChainable) awsstepfunctions.Map
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
