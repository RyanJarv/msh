package voiceidupdatesessionaction

type VoiceIdUpdateSessionAction struct {
    ErrorInfo ErrorInfo `json:"errorInfo,omitempty"`
    Session Session `json:"session,omitempty"`
    SystemAttributes SystemAttributes `json:"systemAttributes,omitempty"`
    Action string `json:"action,omitempty"`
    DomainId string `json:"domainId,omitempty"`
    SourceId string `json:"sourceId,omitempty"`
    Status string `json:"status,omitempty"`
}

func (v *VoiceIdUpdateSessionAction) SetErrorInfo(errorInfo ErrorInfo) {
    v.ErrorInfo = errorInfo
}

func (v *VoiceIdUpdateSessionAction) SetSession(session Session) {
    v.Session = session
}

func (v *VoiceIdUpdateSessionAction) SetSystemAttributes(systemAttributes SystemAttributes) {
    v.SystemAttributes = systemAttributes
}

func (v *VoiceIdUpdateSessionAction) SetAction(action string) {
    v.Action = action
}

func (v *VoiceIdUpdateSessionAction) SetDomainId(domainId string) {
    v.DomainId = domainId
}

func (v *VoiceIdUpdateSessionAction) SetSourceId(sourceId string) {
    v.SourceId = sourceId
}

func (v *VoiceIdUpdateSessionAction) SetStatus(status string) {
    v.Status = status
}
