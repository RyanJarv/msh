package sns

import (
	_ "embed"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/ryanjarv/msh/pkg/utils"
)

func NewSns() (*Sns, error) {
	return &Sns{}, nil
}

type Sns struct{}

func (s *Sns) Name() string { return "sns" }

func (s *Sns) Compile(stack awscdk.Stack, next []interface{}) ([]interface{}, error) {
	subscriptions, ok := utils.EachAs[awssns.ITopicSubscription](next)
	if !ok {
		return nil, fmt.Errorf("expected sns topic, got: %T: %+v", next, next)
	}

	topic := awssns.NewTopic(stack, aws.String(s.Name()), &awssns.TopicProps{})

	for _, sub := range subscriptions {
		topic.AddSubscription(sub)
	}

	target := awseventstargets.NewSnsTopic(topic, &awseventstargets.SnsTopicProps{})

	return []interface{}{target}, nil
}
