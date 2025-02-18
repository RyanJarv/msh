package main

import (
	"flag"
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
	"github.com/ryanjarv/msh/cmd/sfn"
	"github.com/ryanjarv/msh/cmd/sleep"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/pkg/types"
	"strings"
)

var Registry = app.Registry{
	"aws":       aws.New,
	"cron":      cron.New,
	"event":     event.New,
	"exclusive": exclusive.New,
	"filter":    filter.New,
	"foreach":   foreach.New,
	"lambda":    lambda.New,
	"sfn":       sfn.New,
	"sleep":     sleep.New,
	"api":       api.New,
	"build":     build.New,
	"call":      call.New,
	"each":      each.New,
	"mail":      mail.New,
	"sns":       sns.New,
}

func GetStepFunc(cmdName string) func(*app.App, []string) (types.CdkStep, error) {
	flag.Parse()

	cmdName = strings.Trim(cmdName, "@.")

	switch cmdName {
	case "aws":
		return aws.New
	case "cron":
		return cron.New
	case "event":
		return event.New
	case "exclusive":
		return exclusive.New
	case "filter":
		return filter.New
	case "foreach":
		return foreach.New
	case "lambda":
		return lambda.New
	case "sfn":
		return sfn.New
	case "sleep":
		return sleep.New
	case "api":
		return api.New
	case "build":
		return build.New
	case "call":
		return call.New
	case "each":
		return each.New
	case "mail":
		return mail.New
	case "sns":
		return sns.New
	default:
		return nil
	}
}
