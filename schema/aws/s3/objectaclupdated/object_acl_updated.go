package objectaclupdated

type ObjectACLUpdated struct {
    Bucket Bucket `json:"bucket"`
    Object Object `json:"object"`
    RequestId string `json:"request-id"`
    Requester string `json:"requester"`
    SourceIpAddress string `json:"source-ip-address"`
    Version string `json:"version"`
}

func (o *ObjectACLUpdated) SetBucket(bucket Bucket) {
    o.Bucket = bucket
}

func (o *ObjectACLUpdated) SetObject(object Object) {
    o.Object = object
}

func (o *ObjectACLUpdated) SetRequestId(requestId string) {
    o.RequestId = requestId
}

func (o *ObjectACLUpdated) SetRequester(requester string) {
    o.Requester = requester
}

func (o *ObjectACLUpdated) SetSourceIpAddress(sourceIpAddress string) {
    o.SourceIpAddress = sourceIpAddress
}

func (o *ObjectACLUpdated) SetVersion(version string) {
    o.Version = version
}
