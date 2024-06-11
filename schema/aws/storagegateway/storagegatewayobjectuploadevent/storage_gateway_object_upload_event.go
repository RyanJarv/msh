package storagegatewayobjectuploadevent

import (
    "time"
)


type StorageGatewayObjectUploadEvent struct {
    BucketName string `json:"bucket-name"`
    EventType string `json:"event-type"`
    ModificationTime time.Time `json:"modification-time"`
    ObjectKey string `json:"object-key"`
    ObjectSize float64 `json:"object-size"`
    Prefix string `json:"prefix"`
}

func (s *StorageGatewayObjectUploadEvent) SetBucketName(bucketName string) {
    s.BucketName = bucketName
}

func (s *StorageGatewayObjectUploadEvent) SetEventType(eventType string) {
    s.EventType = eventType
}

func (s *StorageGatewayObjectUploadEvent) SetModificationTime(modificationTime time.Time) {
    s.ModificationTime = modificationTime
}

func (s *StorageGatewayObjectUploadEvent) SetObjectKey(objectKey string) {
    s.ObjectKey = objectKey
}

func (s *StorageGatewayObjectUploadEvent) SetObjectSize(objectSize float64) {
    s.ObjectSize = objectSize
}

func (s *StorageGatewayObjectUploadEvent) SetPrefix(prefix string) {
    s.Prefix = prefix
}
