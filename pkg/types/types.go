package types

import (
	"github.com/aws/constructs-go/constructs/v10"
	"io"
)

func NewRegistry(steps ...IStep) Registry {
	r := Registry{}
	for _, step := range steps {
		r[step.GetName()] = step
	}

	return r
}

type Registry map[string]IStep

type IStep interface {
	GetName() string
}

type Step struct {
	Name  string
	Value interface{}
}

func (s Step) GetName() string {
	return s.Name
}

type CdkStep interface {
	Compile(constructs.Construct, interface{}, int) ([]interface{}, error)
}

type CdkStepFanOut interface {
	Compile(constructs.Construct, []interface{}, int) ([]interface{}, error)
}

type Process interface {
	io.ReaderFrom
	io.WriterTo
	Run() error
}

type Deployable interface {
	Deploy() error
}
