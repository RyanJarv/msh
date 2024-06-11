package awsapicallviacloudtrail

type ListTagsForResourceRequestParameters struct {
    ResourceArn string `json:"resource-arn"`
}

func (l *ListTagsForResourceRequestParameters) SetResourceArn(resourceArn string) {
    l.ResourceArn = resourceArn
}
