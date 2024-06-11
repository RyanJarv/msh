package awsapicallviacloudtrail

type TagResourceRequestParameters struct {
    ResourceArn string `json:"resource-arn"`
    Tags Tags `json:"tags"`
}

func (t *TagResourceRequestParameters) SetResourceArn(resourceArn string) {
    t.ResourceArn = resourceArn
}

func (t *TagResourceRequestParameters) SetTags(tags Tags) {
    t.Tags = tags
}
