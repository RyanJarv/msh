package sns

import (
	_ "embed"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func NewSns() (*Sns, error) {
	return &Sns{}, nil
}

type Sns struct{}

func (s *Sns) Name() string { return "sns" }

func (s *Sns) Run(stack awscdk.Stack, next interface{}) (interface{}, error) {
	subscription, ok := next.(awssns.ITopicSubscription)
	if !ok {
		return nil, fmt.Errorf("next step must be a sns topic, got: %T", next)
	}

	topic := awssns.NewTopic(stack, aws.String(s.Name()), &awssns.TopicProps{})
	topic.AddSubscription(subscription)

	return topic, nil
}
