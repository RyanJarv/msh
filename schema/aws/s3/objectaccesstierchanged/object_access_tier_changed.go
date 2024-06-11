package objectaccesstierchanged

type ObjectAccessTierChanged struct {
    Bucket Bucket `json:"bucket"`
    Object Object `json:"object"`
    DestinationAccessTier string `json:"destination-access-tier"`
    RequestId string `json:"request-id"`
    Requester string `json:"requester"`
    Version string `json:"version"`
}

func (o *ObjectAccessTierChanged) SetBucket(bucket Bucket) {
    o.Bucket = bucket
}

func (o *ObjectAccessTierChanged) SetObject(object Object) {
    o.Object = object
}

func (o *ObjectAccessTierChanged) SetDestinationAccessTier(destinationAccessTier string) {
    o.DestinationAccessTier = destinationAccessTier
}

func (o *ObjectAccessTierChanged) SetRequestId(requestId string) {
    o.RequestId = requestId
}

func (o *ObjectAccessTierChanged) SetRequester(requester string) {
    o.Requester = requester
}

func (o *ObjectAccessTierChanged) SetVersion(version string) {
    o.Version = version
}
