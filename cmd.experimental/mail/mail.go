package mail

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/ryanjarv/msh/pkg/app"
)

var opts = struct {
	To      *string
	Subject *string
}{
	To:      flag.String("to", "", "Address to send notifications to."),
	Subject: flag.String("subject", "", "Subject of the notification email."),
}

func New(app app.App) (*Mail, error) {
	if *opts.To == "" {
		return nil, fmt.Errorf("missing required argument: to")
	}

	if *opts.Subject == "" {
		return nil, fmt.Errorf("missing required argument: to")
	}
	return &Mail{
		To:      opts.To,
		Subject: opts.Subject,
	}, nil
}

type Mail struct {
	To      *string
	Subject *string
	awssns.ITopicSubscription
}

func (s Mail) GetName() string { return "mail" }

func (s *Mail) Compile(stack constructs.Construct, i int) error {
	s.ITopicSubscription = awssnssubscriptions.NewEmailSubscription(s.To, &awssnssubscriptions.EmailSubscriptionProps{})
	return nil
}
