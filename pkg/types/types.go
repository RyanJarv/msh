package types

import (
	"fmt"
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

		out := reflect.New(f.Out(0).Elem())
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
	Compile(construct constructs.Construct, i int) error
}

type IChain interface {
	awsstepfunctions.INextable
	awsstepfunctions.IChainable
}

type IIterator interface {
	Iterator(iterator awsstepfunctions.IChainable) awsstepfunctions.Map
}
