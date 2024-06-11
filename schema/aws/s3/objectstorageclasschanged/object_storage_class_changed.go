package objectstorageclasschanged

type ObjectStorageClassChanged struct {
    Bucket Bucket `json:"bucket"`
    Object Object `json:"object"`
    DestinationStorageClass string `json:"destination-storage-class"`
    RequestId string `json:"request-id"`
    Requester string `json:"requester"`
    Version string `json:"version"`
}

func (o *ObjectStorageClassChanged) SetBucket(bucket Bucket) {
    o.Bucket = bucket
}

func (o *ObjectStorageClassChanged) SetObject(object Object) {
    o.Object = object
}

func (o *ObjectStorageClassChanged) SetDestinationStorageClass(destinationStorageClass string) {
    o.DestinationStorageClass = destinationStorageClass
}

func (o *ObjectStorageClassChanged) SetRequestId(requestId string) {
    o.RequestId = requestId
}

func (o *ObjectStorageClassChanged) SetRequester(requester string) {
    o.Requester = requester
}

func (o *ObjectStorageClassChanged) SetVersion(version string) {
    o.Version = version
}
