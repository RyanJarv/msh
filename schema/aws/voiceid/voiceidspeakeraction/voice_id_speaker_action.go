package voiceidspeakeraction

type VoiceIdSpeakerAction struct {
    Data Data `json:"data"`
    ErrorInfo ErrorInfo `json:"errorInfo"`
    SystemAttributes SystemAttributes `json:"systemAttributes"`
    Action string `json:"action"`
    DomainId string `json:"domainId"`
    GeneratedSpeakerId string `json:"generatedSpeakerId"`
    SourceId string `json:"sourceId"`
    Status string `json:"status"`
}

func (v *VoiceIdSpeakerAction) SetData(data Data) {
    v.Data = data
}

func (v *VoiceIdSpeakerAction) SetErrorInfo(errorInfo ErrorInfo) {
    v.ErrorInfo = errorInfo
}

func (v *VoiceIdSpeakerAction) SetSystemAttributes(systemAttributes SystemAttributes) {
    v.SystemAttributes = systemAttributes
}

func (v *VoiceIdSpeakerAction) SetAction(action string) {
    v.Action = action
}

func (v *VoiceIdSpeakerAction) SetDomainId(domainId string) {
    v.DomainId = domainId
}

func (v *VoiceIdSpeakerAction) SetGeneratedSpeakerId(generatedSpeakerId string) {
    v.GeneratedSpeakerId = generatedSpeakerId
}

func (v *VoiceIdSpeakerAction) SetSourceId(sourceId string) {
    v.SourceId = sourceId
}

func (v *VoiceIdSpeakerAction) SetStatus(status string) {
    v.Status = status
}
