package awsapicallviacloudtrail

type UntagResourceRequestParameters struct {
    ResourceArn string `json:"resource-arn"`
    TagKeys string `json:"tagKeys"`
}

func (u *UntagResourceRequestParameters) SetResourceArn(resourceArn string) {
    u.ResourceArn = resourceArn
}

func (u *UntagResourceRequestParameters) SetTagKeys(tagKeys string) {
    u.TagKeys = tagKeys
}
