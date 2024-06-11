package objectrestoreinitiated

type ObjectRestoreInitiated struct {
    Bucket Bucket `json:"bucket"`
    Object Object `json:"object"`
    RequestId string `json:"request-id"`
    Requester string `json:"requester"`
    SourceIpAddress string `json:"source-ip-address"`
    SourceStorageClass string `json:"source-storage-class"`
    Version string `json:"version"`
}

func (o *ObjectRestoreInitiated) SetBucket(bucket Bucket) {
    o.Bucket = bucket
}

func (o *ObjectRestoreInitiated) SetObject(object Object) {
    o.Object = object
}

func (o *ObjectRestoreInitiated) SetRequestId(requestId string) {
    o.RequestId = requestId
}

func (o *ObjectRestoreInitiated) SetRequester(requester string) {
    o.Requester = requester
}

func (o *ObjectRestoreInitiated) SetSourceIpAddress(sourceIpAddress string) {
    o.SourceIpAddress = sourceIpAddress
}

func (o *ObjectRestoreInitiated) SetSourceStorageClass(sourceStorageClass string) {
    o.SourceStorageClass = sourceStorageClass
}

func (o *ObjectRestoreInitiated) SetVersion(version string) {
    o.Version = version
}
