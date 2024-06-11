package objecttagsdeleted

type ObjectTagsDeleted struct {
    Bucket Bucket `json:"bucket"`
    Object Object `json:"object"`
    RequestId string `json:"request-id"`
    Requester string `json:"requester"`
    SourceIpAddress string `json:"source-ip-address"`
    Version string `json:"version"`
}

func (o *ObjectTagsDeleted) SetBucket(bucket Bucket) {
    o.Bucket = bucket
}

func (o *ObjectTagsDeleted) SetObject(object Object) {
    o.Object = object
}

func (o *ObjectTagsDeleted) SetRequestId(requestId string) {
    o.RequestId = requestId
}

func (o *ObjectTagsDeleted) SetRequester(requester string) {
    o.Requester = requester
}

func (o *ObjectTagsDeleted) SetSourceIpAddress(sourceIpAddress string) {
    o.SourceIpAddress = sourceIpAddress
}

func (o *ObjectTagsDeleted) SetVersion(version string) {
    o.Version = version
}
