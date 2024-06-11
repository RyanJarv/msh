package objectrestorecompleted

import (
    "time"
)


type ObjectRestoreCompleted struct {
    Bucket Bucket `json:"bucket"`
    Object Object `json:"object"`
    RequestId string `json:"request-id"`
    Requester string `json:"requester"`
    RestoreExpiryTime time.Time `json:"restore-expiry-time,omitempty"`
    SourceStorageClass string `json:"source-storage-class"`
    Version string `json:"version"`
}

func (o *ObjectRestoreCompleted) SetBucket(bucket Bucket) {
    o.Bucket = bucket
}

func (o *ObjectRestoreCompleted) SetObject(object Object) {
    o.Object = object
}

func (o *ObjectRestoreCompleted) SetRequestId(requestId string) {
    o.RequestId = requestId
}

func (o *ObjectRestoreCompleted) SetRequester(requester string) {
    o.Requester = requester
}

func (o *ObjectRestoreCompleted) SetRestoreExpiryTime(restoreExpiryTime time.Time) {
    o.RestoreExpiryTime = restoreExpiryTime
}

func (o *ObjectRestoreCompleted) SetSourceStorageClass(sourceStorageClass string) {
    o.SourceStorageClass = sourceStorageClass
}

func (o *ObjectRestoreCompleted) SetVersion(version string) {
    o.Version = version
}
