package fd

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/samber/lo"
	"strings"
)

func SetupSqsFd(cfg aws.Config, name, fd string) (arn string, uri string) {
	L.Debug.Println("creating sqs queue:", name, fd)
	client := sqs.NewFromConfig(cfg)

	queueName := fmt.Sprintf("%s-%s", name, fd)
	resp := lo.Must(client.CreateQueue(context.TODO(), &sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
	}))
	L.Debug.Println("created sqs queue:", *resp.QueueUrl)

	// https://sqs.us-east-1.amazonaws.com/123456789012/cat-msh-0a0c3f9d7adb3006a923c4fd8884951b068f44b5-stdout3
	p := strings.Split(*resp.QueueUrl, "/")
	domain := p[2]
	account := p[3]
	region := strings.Split(domain, ".")[1]

	return fmt.Sprintf("arn:aws:sqs:%s:%s:%s", region, account, queueName), *resp.QueueUrl
}
