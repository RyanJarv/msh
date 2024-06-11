package awsapicallviacloudtrail

type ResponseElementsItem struct {
    InstanceId string `json:"instanceId"`
}

func (r *ResponseElementsItem) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
