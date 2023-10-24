package mail

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-sdk-go-v2/aws"
)

var args = struct {
	To      *string
	Subject *string
}{
	To:      flag.String("to", "", "Address to send notifications to."),
	Subject: flag.String("subject", "", "Subject of the notification email."),
}

func NewMail() (*Mail, error) {
	if *args.To == "" {
		return nil, fmt.Errorf("missing required argument: to")
	}

	if *args.Subject == "" {
		return nil, fmt.Errorf("missing required argument: to")
	}
	return &Mail{
		To:      args.To,
		Subject: args.Subject,
	}, nil
}

type Mail struct {
	To      *string
	Subject *string
}

func (s *Mail) Name() string { return "mail" }

func (s *Mail) Run(stack awscdk.Stack, last interface{}) (interface{}, error) {
	topic, ok := last.(awssns.Topic)
	if !ok {
		return nil, fmt.Errorf("last step must be a sns topic rule, got: %T", last)
	}

	awssns.NewSubscription(stack, aws.String(s.Name()), &awssns.SubscriptionProps{
		Endpoint: s.To,
		Protocol: awssns.SubscriptionProtocol_EMAIL,
		Topic:    topic,

		// TODO: implement these
		//DeadLetterQueue:             nil,
		//FilterPolicy:                nil,
		//FilterPolicyWithMessageBody: nil,
	})

	// Can't chain anything after this
	return nil, nil
}
