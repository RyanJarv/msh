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
	"github.com/ryanjarv/msh/cmd/exclusive"
	"github.com/ryanjarv/msh/cmd/filter"
	"github.com/ryanjarv/msh/cmd/foreach"
	"github.com/ryanjarv/msh/cmd/lambda"
	"github.com/ryanjarv/msh/cmd/name"
	"github.com/ryanjarv/msh/cmd/sfn"
	"github.com/ryanjarv/msh/cmd/sleep"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/types"
	"github.com/ryanjarv/msh/pkg/utils"
	"strings"
)

var Registry = types.NewRegistry(
	aws.New,
	cron.New,
	event.New,
	exclusive.New,
	filter.New,
	foreach.New,
	items.New,
	lambda.New,
	name.New,
	sfn.New,
	sleep.New,
	api.New,
	build.New,
	call.New,
	each.New,
	mail.New,
	sns.New,
)

func NewModule(app *app.App, cmdName string) (step types.CdkStep, err error) {
	flag.Parse()

	cmdName = strings.Trim(cmdName, "@.")

	switch cmdName {
	case "aws":
		step, err = aws.New(app)
	case "cron":
		step, err = cron.New(app)
	case "event":
		step, err = event.New(app)
	case "exclusive":
		step, err = exclusive.New(app)
	case "filter":
		step, err = filter.New(app)
	case "foreach":
		step, err = foreach.New(app)
	case "items":
		step, err = items.New(app)
	case "lambda":
		step, err = lambda.New(app)
	case "name":
		step, err = name.New(app)
	case "sfn":
		step, err = sfn.New(app)
	case "sleep":
		step, err = sleep.New(app)
	case "api":
		step, err = api.New(app)
	case "build":
		step, err = build.New(app)
	case "call":
		step, err = call.New(app)
	case "each":
		step, err = each.New(app)
	case "mail":
		step, err = mail.New(app)
	case "sns":
		step, err = sns.New(app)
	default:
		err = fmt.Errorf("unknown module: %s", cmdName)
	}

	return step, utils.Wrap(err, cmdName)
}
