package mediapackagekeyprovidernotification

type MediaPackageKeyProviderNotification struct {
    Event string `json:"event"`
    Message string `json:"message"`
}

func (m *MediaPackageKeyProviderNotification) SetEvent(event string) {
    m.Event = event
}

func (m *MediaPackageKeyProviderNotification) SetMessage(message string) {
    m.Message = message
}
