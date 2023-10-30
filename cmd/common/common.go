package common

import (
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

var Registry = types.Registry{
	"sfn":    &sfn.Sfn{},
	"event":  &event.Event{},
	"lambda": &lambda.LambdaCmd{},
	"sleep":  &sleep.SleepCmd{},
	"mail":   &mail.Mail{},
	"sns":    &sns.Sns{},
	"each":   &each.Each{},
	"filter": &sfn_filter.Filter{},
	"map":    &sfn_map.Map{},
}
