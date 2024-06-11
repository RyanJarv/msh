package objectcreated

type ObjectCreated struct {
    Bucket Bucket `json:"bucket"`
    Object Object `json:"object"`
    Reason string `json:"reason"`
    RequestId string `json:"request-id"`
    Requester string `json:"requester"`
    SourceIpAddress string `json:"source-ip-address"`
    Version string `json:"version"`
}

func (o *ObjectCreated) SetBucket(bucket Bucket) {
    o.Bucket = bucket
}

func (o *ObjectCreated) SetObject(object Object) {
    o.Object = object
}

func (o *ObjectCreated) SetReason(reason string) {
    o.Reason = reason
}

func (o *ObjectCreated) SetRequestId(requestId string) {
    o.RequestId = requestId
}

func (o *ObjectCreated) SetRequester(requester string) {
    o.Requester = requester
}

func (o *ObjectCreated) SetSourceIpAddress(sourceIpAddress string) {
    o.SourceIpAddress = sourceIpAddress
}

func (o *ObjectCreated) SetVersion(version string) {
    o.Version = version
}
