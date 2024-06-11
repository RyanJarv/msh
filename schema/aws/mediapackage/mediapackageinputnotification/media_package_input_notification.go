package mediapackageinputnotification

type MediaPackageInputNotification struct {
    Event string `json:"event"`
    Message string `json:"message"`
    PackagingConfigurationId string `json:"packaging_configuration_id,omitempty"`
    ManifestUrls []string `json:"manifest_urls,omitempty"`
}

func (m *MediaPackageInputNotification) SetEvent(event string) {
    m.Event = event
}

func (m *MediaPackageInputNotification) SetMessage(message string) {
    m.Message = message
}

func (m *MediaPackageInputNotification) SetPackagingConfigurationId(packagingConfigurationId string) {
    m.PackagingConfigurationId = packagingConfigurationId
}

func (m *MediaPackageInputNotification) SetManifestUrls(manifestUrls []string) {
    m.ManifestUrls = manifestUrls
}
