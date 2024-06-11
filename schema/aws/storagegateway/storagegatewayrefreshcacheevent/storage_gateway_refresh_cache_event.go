package storagegatewayrefreshcacheevent

import (
    "time"
)


type StorageGatewayRefreshCacheEvent struct {
    Completed time.Time `json:"completed"`
    EventType string `json:"event-type"`
    FolderList []string `json:"folderList"`
    NotificationId string `json:"notification-id"`
    Started time.Time `json:"started"`
}

func (s *StorageGatewayRefreshCacheEvent) SetCompleted(completed time.Time) {
    s.Completed = completed
}

func (s *StorageGatewayRefreshCacheEvent) SetEventType(eventType string) {
    s.EventType = eventType
}

func (s *StorageGatewayRefreshCacheEvent) SetFolderList(folderList []string) {
    s.FolderList = folderList
}

func (s *StorageGatewayRefreshCacheEvent) SetNotificationId(notificationId string) {
    s.NotificationId = notificationId
}

func (s *StorageGatewayRefreshCacheEvent) SetStarted(started time.Time) {
    s.Started = started
}
