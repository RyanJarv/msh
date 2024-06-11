package voiceidevaluatesessionaction

type VoiceIdEvaluateSessionAction struct {
    ErrorInfo ErrorInfo `json:"errorInfo,omitempty"`
    Session Session `json:"session,omitempty"`
    SystemAttributes SystemAttributes `json:"systemAttributes,omitempty"`
    Action string `json:"action,omitempty"`
    DomainId string `json:"domainId,omitempty"`
    SourceId string `json:"sourceId,omitempty"`
    Status string `json:"status,omitempty"`
}

func (v *VoiceIdEvaluateSessionAction) SetErrorInfo(errorInfo ErrorInfo) {
    v.ErrorInfo = errorInfo
}

func (v *VoiceIdEvaluateSessionAction) SetSession(session Session) {
    v.Session = session
}

func (v *VoiceIdEvaluateSessionAction) SetSystemAttributes(systemAttributes SystemAttributes) {
    v.SystemAttributes = systemAttributes
}

func (v *VoiceIdEvaluateSessionAction) SetAction(action string) {
    v.Action = action
}

func (v *VoiceIdEvaluateSessionAction) SetDomainId(domainId string) {
    v.DomainId = domainId
}

func (v *VoiceIdEvaluateSessionAction) SetSourceId(sourceId string) {
    v.SourceId = sourceId
}

func (v *VoiceIdEvaluateSessionAction) SetStatus(status string) {
    v.Status = status
}
