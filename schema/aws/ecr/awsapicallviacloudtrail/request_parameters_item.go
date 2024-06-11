package awsapicallviacloudtrail

type RequestParametersItem struct {
    ImageTag string `json:"imageTag"`
}

func (r *RequestParametersItem) SetImageTag(imageTag string) {
    r.ImageTag = imageTag
}
