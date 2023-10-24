package app

import (
	"github.com/ryanjarv/msh/cmd/event"
	"github.com/ryanjarv/msh/cmd/lambda"
	"github.com/ryanjarv/msh/cmd/mail"
	"github.com/ryanjarv/msh/cmd/sfn"
	"github.com/ryanjarv/msh/cmd/sleep"
	"github.com/ryanjarv/msh/cmd/sns"
	"github.com/ryanjarv/msh/pkg/types"
)

// Registry maintains the build lookup map for construct types.
//
// It is used during the build process to resolve each step type in the pipeline's app.
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
	r.Add(&mail.Mail{})
	r.Add(&sns.Sns{})

	return r
}
