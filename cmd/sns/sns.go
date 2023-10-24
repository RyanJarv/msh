package sns

import (
	_ "embed"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func NewSns() (*Sns, error) {
	return &Sns{}, nil
}

type Sns struct{}

func (s *Sns) Name() string { return "sns" }

func (s *Sns) Run(stack awscdk.Stack, last interface{}) (interface{}, error) {
	rule, ok := last.(awsevents.Rule)
	if !ok {
		return nil, fmt.Errorf("last step must be eventbridge rule, got: %T", last)
	}

	topic := awssns.NewTopic(stack, aws.String(s.Name()), &awssns.TopicProps{})

	target := awseventstargets.NewSnsTopic(topic, &awseventstargets.SnsTopicProps{
		// TODO: implement these
		//MaxEventAge: nil,
		//RetryAttempts: nil,
		//DeadLetterQueue: nil,
		//Message: nil,
	})

	rule.AddTarget(target)

	return topic, nil
}
