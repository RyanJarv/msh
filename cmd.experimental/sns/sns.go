package sns

import (
	_ "embed"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/ryanjarv/msh/pkg/app"
)

func New(app *app.App) (*Sns, error) {
	return &Sns{}, nil
}

type Sns struct {
	awssns.Topic
	awseventstargets.SnsTopic
}

func (s Sns) GetName() string { return "sns" }

func (s Sns) Init(scope constructs.Construct, i int) error {
	s.Topic = awssns.NewTopic(scope, aws.String(s.GetName()), &awssns.TopicProps{})
	s.SnsTopic = awseventstargets.NewSnsTopic(s.Topic, &awseventstargets.SnsTopicProps{})
	return nil
}
