package voiceidsessionspeakerenrollmentaction

type VoiceIdSessionSpeakerEnrollmentAction struct {
    ErrorInfo ErrorInfo `json:"errorInfo"`
    SystemAttributes SystemAttributes `json:"systemAttributes"`
    Action string `json:"action"`
    DomainId string `json:"domainId"`
    SessionId string `json:"sessionId"`
    SessionName string `json:"sessionName"`
    SourceId string `json:"sourceId"`
    Status string `json:"status"`
}

func (v *VoiceIdSessionSpeakerEnrollmentAction) SetErrorInfo(errorInfo ErrorInfo) {
    v.ErrorInfo = errorInfo
}

func (v *VoiceIdSessionSpeakerEnrollmentAction) SetSystemAttributes(systemAttributes SystemAttributes) {
    v.SystemAttributes = systemAttributes
}

func (v *VoiceIdSessionSpeakerEnrollmentAction) SetAction(action string) {
    v.Action = action
}

func (v *VoiceIdSessionSpeakerEnrollmentAction) SetDomainId(domainId string) {
    v.DomainId = domainId
}

func (v *VoiceIdSessionSpeakerEnrollmentAction) SetSessionId(sessionId string) {
    v.SessionId = sessionId
}

func (v *VoiceIdSessionSpeakerEnrollmentAction) SetSessionName(sessionName string) {
    v.SessionName = sessionName
}

func (v *VoiceIdSessionSpeakerEnrollmentAction) SetSourceId(sourceId string) {
    v.SourceId = sourceId
}

func (v *VoiceIdSessionSpeakerEnrollmentAction) SetStatus(status string) {
    v.Status = status
}
