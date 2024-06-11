package voiceidstartsessionaction

type VoiceIdStartSessionAction struct {
    ErrorInfo ErrorInfo `json:"errorInfo,omitempty"`
    Session Session `json:"session,omitempty"`
    SystemAttributes SystemAttributes `json:"systemAttributes,omitempty"`
    Action string `json:"action,omitempty"`
    DomainId string `json:"domainId,omitempty"`
    SourceId string `json:"sourceId,omitempty"`
    Status string `json:"status,omitempty"`
}

func (v *VoiceIdStartSessionAction) SetErrorInfo(errorInfo ErrorInfo) {
    v.ErrorInfo = errorInfo
}

func (v *VoiceIdStartSessionAction) SetSession(session Session) {
    v.Session = session
}

func (v *VoiceIdStartSessionAction) SetSystemAttributes(systemAttributes SystemAttributes) {
    v.SystemAttributes = systemAttributes
}

func (v *VoiceIdStartSessionAction) SetAction(action string) {
    v.Action = action
}

func (v *VoiceIdStartSessionAction) SetDomainId(domainId string) {
    v.DomainId = domainId
}

func (v *VoiceIdStartSessionAction) SetSourceId(sourceId string) {
    v.SourceId = sourceId
}

func (v *VoiceIdStartSessionAction) SetStatus(status string) {
    v.Status = status
}
