package voiceidbatchspeakerenrollmentaction

type VoiceIdBatchSpeakerEnrollmentAction struct {
    Data Data `json:"data,omitempty"`
    ErrorInfo ErrorInfo `json:"errorInfo,omitempty"`
    Action string `json:"action,omitempty"`
    BatchJobId string `json:"batchJobId,omitempty"`
    DomainId string `json:"domainId,omitempty"`
    SourceId string `json:"sourceId,omitempty"`
    Status string `json:"status,omitempty"`
}

func (v *VoiceIdBatchSpeakerEnrollmentAction) SetData(data Data) {
    v.Data = data
}

func (v *VoiceIdBatchSpeakerEnrollmentAction) SetErrorInfo(errorInfo ErrorInfo) {
    v.ErrorInfo = errorInfo
}

func (v *VoiceIdBatchSpeakerEnrollmentAction) SetAction(action string) {
    v.Action = action
}

func (v *VoiceIdBatchSpeakerEnrollmentAction) SetBatchJobId(batchJobId string) {
    v.BatchJobId = batchJobId
}

func (v *VoiceIdBatchSpeakerEnrollmentAction) SetDomainId(domainId string) {
    v.DomainId = domainId
}

func (v *VoiceIdBatchSpeakerEnrollmentAction) SetSourceId(sourceId string) {
    v.SourceId = sourceId
}

func (v *VoiceIdBatchSpeakerEnrollmentAction) SetStatus(status string) {
    v.Status = status
}
