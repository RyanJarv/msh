package voiceidfraudsteraction

type VoiceIdFraudsterAction struct {
    Data Data `json:"data,omitempty"`
    ErrorInfo ErrorInfo `json:"errorInfo,omitempty"`
    Action string `json:"action,omitempty"`
    DomainId string `json:"domainId,omitempty"`
    GeneratedFraudsterId interface{} `json:"generatedFraudsterId,omitempty"`
    SourceId string `json:"sourceId,omitempty"`
    Status string `json:"status,omitempty"`
    WatchlistIds []string `json:"watchlistIds,omitempty"`
}

func (v *VoiceIdFraudsterAction) SetData(data Data) {
    v.Data = data
}

func (v *VoiceIdFraudsterAction) SetErrorInfo(errorInfo ErrorInfo) {
    v.ErrorInfo = errorInfo
}

func (v *VoiceIdFraudsterAction) SetAction(action string) {
    v.Action = action
}

func (v *VoiceIdFraudsterAction) SetDomainId(domainId string) {
    v.DomainId = domainId
}

func (v *VoiceIdFraudsterAction) SetGeneratedFraudsterId(generatedFraudsterId interface{}) {
    v.GeneratedFraudsterId = generatedFraudsterId
}

func (v *VoiceIdFraudsterAction) SetSourceId(sourceId string) {
    v.SourceId = sourceId
}

func (v *VoiceIdFraudsterAction) SetStatus(status string) {
    v.Status = status
}

func (v *VoiceIdFraudsterAction) SetWatchlistIds(watchlistIds []string) {
    v.WatchlistIds = watchlistIds
}
