package objectdeleted

type ObjectDeleted struct {
    Bucket Bucket `json:"bucket"`
    Object Object `json:"object"`
    DeletionType string `json:"deletion-type"`
    Reason string `json:"reason"`
    RequestId string `json:"request-id"`
    Requester string `json:"requester"`
    SourceIpAddress string `json:"source-ip-address"`
    Version string `json:"version"`
}

func (o *ObjectDeleted) SetBucket(bucket Bucket) {
    o.Bucket = bucket
}

func (o *ObjectDeleted) SetObject(object Object) {
    o.Object = object
}

func (o *ObjectDeleted) SetDeletionType(deletionType string) {
    o.DeletionType = deletionType
}

func (o *ObjectDeleted) SetReason(reason string) {
    o.Reason = reason
}

func (o *ObjectDeleted) SetRequestId(requestId string) {
    o.RequestId = requestId
}

func (o *ObjectDeleted) SetRequester(requester string) {
    o.Requester = requester
}

func (o *ObjectDeleted) SetSourceIpAddress(sourceIpAddress string) {
    o.SourceIpAddress = sourceIpAddress
}

func (o *ObjectDeleted) SetVersion(version string) {
    o.Version = version
}
