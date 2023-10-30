package common

import (
	"github.com/ryanjarv/msh/cmd/aws"
	"github.com/ryanjarv/msh/cmd/each"
	"github.com/ryanjarv/msh/cmd/event"
	"github.com/ryanjarv/msh/cmd/lambda"
	"github.com/ryanjarv/msh/cmd/mail"
	"github.com/ryanjarv/msh/cmd/sfn"
	sfn_filter "github.com/ryanjarv/msh/cmd/sfn.filter"
	sfn_map "github.com/ryanjarv/msh/cmd/sfn.map"
	"github.com/ryanjarv/msh/cmd/sleep"
	"github.com/ryanjarv/msh/cmd/sns"
	"github.com/ryanjarv/msh/pkg/types"
)

var Registry = types.NewRegistry(
	&sfn.Sfn{},
	&event.Event{},
	&lambda.Lambda{},
	&aws.AwsCmd{},
	&sleep.SleepCmd{},
	&mail.Mail{},
	&sns.Sns{},
	&each.Each{},
	&sfn_filter.Filter{},
	&sfn_map.Map{},
)
