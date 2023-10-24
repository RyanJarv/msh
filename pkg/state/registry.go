package state

import (
	"github.com/ryanjarv/msh/pkg/providers/event"
	"github.com/ryanjarv/msh/pkg/providers/lambda"
	"github.com/ryanjarv/msh/pkg/providers/sfn"
	"github.com/ryanjarv/msh/pkg/providers/sleep"
	"github.com/ryanjarv/msh/pkg/types"
)

// Registry maintains the build lookup map for construct types.
//
// It is used during the build process to resolve each step type in the pipeline's state.
type Registry map[string]types.IStep

func (s *Registry) Add(c types.IStep) {
	(*s)[c.Name()] = c
}

func SetupRegistry() Registry {
	var r = Registry{}

	r.Add(&sfn.Sfn{})
	r.Add(&event.Event{})
	r.Add(&lambda.LambdaCmd{})
	r.Add(&sleep.SleepCmd{})

	return r
}
