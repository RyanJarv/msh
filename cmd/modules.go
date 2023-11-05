
package main

import (
	"flag"
	"fmt"
	"github.com/ryanjarv/msh/pkg/types"
	"github.com/ryanjarv/msh/pkg/app"
	"github.com/ryanjarv/msh/cmd/api"
	"github.com/ryanjarv/msh/cmd/aws"
	"github.com/ryanjarv/msh/cmd/build"
	"github.com/ryanjarv/msh/cmd/call"
	"github.com/ryanjarv/msh/cmd/cron"
	"github.com/ryanjarv/msh/cmd/each"
	"github.com/ryanjarv/msh/cmd/event"
	"github.com/ryanjarv/msh/cmd/lambda"
	"github.com/ryanjarv/msh/cmd/mail"
	"github.com/ryanjarv/msh/cmd/sfn"
	"github.com/ryanjarv/msh/cmd/sfn_filter"
	"github.com/ryanjarv/msh/cmd/sfn_map"
	"github.com/ryanjarv/msh/cmd/sleep"
	"github.com/ryanjarv/msh/cmd/sns"
)

var Registry = types.NewRegistry(
	api.New,
	aws.New,
	build.New,
	call.New,
	cron.New,
	each.New,
	event.New,
	lambda.New,
	mail.New,
	sfn.New,
	sfn_filter.New,
	sfn_map.New,
	sleep.New,
	sns.New,
)

func NewModule(app app.App, name string) (types.CdkStep, error) {
    flag.Parse()

	switch name {
	case ".api":
		 return api.New(app)
	case ".aws":
		 return aws.New(app)
	case ".build":
		 return build.New(app)
	case ".call":
		 return call.New(app)
	case ".cron":
		 return cron.New(app)
	case ".each":
		 return each.New(app)
	case ".event":
		 return event.New(app)
	case ".lambda":
		 return lambda.New(app)
	case ".mail":
		 return mail.New(app)
	case ".sfn":
		 return sfn.New(app)
	case ".sfn_filter":
		 return sfn_filter.New(app)
	case ".sfn_map":
		 return sfn_map.New(app)
	case ".sleep":
		 return sleep.New(app)
	case ".sns":
		 return sns.New(app)
	}

	return nil, fmt.Errorf("unknown module: %s", name)
}
