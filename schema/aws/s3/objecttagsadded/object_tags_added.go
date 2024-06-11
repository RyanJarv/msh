package objecttagsadded

type ObjectTagsAdded struct {
    Bucket Bucket `json:"bucket"`
    Object Object `json:"object"`
    RequestId string `json:"request-id"`
    Requester string `json:"requester"`
    SourceIpAddress string `json:"source-ip-address"`
    Version string `json:"version"`
}

func (o *ObjectTagsAdded) SetBucket(bucket Bucket) {
    o.Bucket = bucket
}

func (o *ObjectTagsAdded) SetObject(object Object) {
    o.Object = object
}

func (o *ObjectTagsAdded) SetRequestId(requestId string) {
    o.RequestId = requestId
}

func (o *ObjectTagsAdded) SetRequester(requester string) {
    o.Requester = requester
}

func (o *ObjectTagsAdded) SetSourceIpAddress(sourceIpAddress string) {
    o.SourceIpAddress = sourceIpAddress
}

func (o *ObjectTagsAdded) SetVersion(version string) {
    o.Version = version
}
