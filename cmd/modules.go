package main

import (
	"flag"
	"fmt"
	"github.com/ryanjarv/msh/cmd.experimental/api"
	"github.com/ryanjarv/msh/cmd.experimental/build"
	"github.com/ryanjarv/msh/cmd.experimental/call"
	"github.com/ryanjarv/msh/cmd.experimental/each"
	"github.com/ryanjarv/msh/cmd.experimental/mail"
	"github.com/ryanjarv/msh/cmd.experimental/sns"
	"github.com/ryanjarv/msh/cmd/aws"
	"github.com/ryanjarv/msh/cmd/cron"
	"github.com/ryanjarv/msh/cmd/event"
	"github.com/ryanjarv/msh/cmd/filter"
	"github.com/ryanjarv/msh/cmd/foreach"
	"github.com/ryanjarv/msh/cmd/lambda"
	"github.com/ryanjarv/msh/cmd/sfn"
	"github.com/ryanjarv/msh/cmd/sleep"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/types"
)

var Registry = types.NewRegistry(
	aws.New,
	cron.New,
	event.New,
	filter.New,
	foreach.New,
	lambda.New,
	sfn.New,
	sleep.New,
	api.New,
	build.New,
	call.New,
	each.New,
	mail.New,
	sns.New,
)

func NewModule(app app.App, name string) (types.CdkStep, error) {
	flag.Parse()

	switch name {
	case ".{cmd/aws aws}":
		return aws.New(app)
	case ".{cmd/cron cron}":
		return cron.New(app)
	case ".{cmd/event event}":
		return event.New(app)
	case ".{cmd/filter filter}":
		return filter.New(app)
	case ".{cmd/foreach foreach}":
		return foreach.New(app)
	case ".{cmd/lambda lambda}":
		return lambda.New(app)
	case ".{cmd/sfn sfn}":
		return sfn.New(app)
	case ".{cmd/sleep sleep}":
		return sleep.New(app)
	case ".{cmd.experimental/api api}":
		return api.New(app)
	case ".{cmd.experimental/build build}":
		return build.New(app)
	case ".{cmd.experimental/call call}":
		return call.New(app)
	case ".{cmd.experimental/each each}":
		return each.New(app)
	case ".{cmd.experimental/mail mail}":
		return mail.New(app)
	case ".{cmd.experimental/sns sns}":
		return sns.New(app)
	}

	return nil, fmt.Errorf("unknown module: %s", name)
}
