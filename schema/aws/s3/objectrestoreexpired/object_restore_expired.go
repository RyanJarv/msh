package objectrestoreexpired

type ObjectRestoreExpired struct {
    Bucket Bucket `json:"bucket"`
    Object Object `json:"object"`
    RequestId string `json:"request-id"`
    Requester string `json:"requester"`
    Version string `json:"version"`
}

func (o *ObjectRestoreExpired) SetBucket(bucket Bucket) {
    o.Bucket = bucket
}

func (o *ObjectRestoreExpired) SetObject(object Object) {
    o.Object = object
}

func (o *ObjectRestoreExpired) SetRequestId(requestId string) {
    o.RequestId = requestId
}

func (o *ObjectRestoreExpired) SetRequester(requester string) {
    o.Requester = requester
}

func (o *ObjectRestoreExpired) SetVersion(version string) {
    o.Version = version
}
