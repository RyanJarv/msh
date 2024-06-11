package storagegatewayfileuploadevent

import (
    "time"
)


type StorageGatewayFileUploadEvent struct {
    Completed time.Time `json:"completed"`
    EventType string `json:"event-type"`
    NotificationId string `json:"notification-id"`
    RequestReceived time.Time `json:"request-received"`
}

func (s *StorageGatewayFileUploadEvent) SetCompleted(completed time.Time) {
    s.Completed = completed
}

func (s *StorageGatewayFileUploadEvent) SetEventType(eventType string) {
    s.EventType = eventType
}

func (s *StorageGatewayFileUploadEvent) SetNotificationId(notificationId string) {
    s.NotificationId = notificationId
}

func (s *StorageGatewayFileUploadEvent) SetRequestReceived(requestReceived time.Time) {
    s.RequestReceived = requestReceived
}
