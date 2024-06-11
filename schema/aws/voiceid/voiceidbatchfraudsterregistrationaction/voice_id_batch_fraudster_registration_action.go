package voiceidbatchfraudsterregistrationaction

type VoiceIdBatchFraudsterRegistrationAction struct {
    Data Data `json:"data,omitempty"`
    ErrorInfo ErrorInfo `json:"errorInfo,omitempty"`
    Action string `json:"action,omitempty"`
    BatchJobId string `json:"batchJobId,omitempty"`
    DomainId string `json:"domainId,omitempty"`
    SourceId string `json:"sourceId,omitempty"`
    Status string `json:"status,omitempty"`
}

func (v *VoiceIdBatchFraudsterRegistrationAction) SetData(data Data) {
    v.Data = data
}

func (v *VoiceIdBatchFraudsterRegistrationAction) SetErrorInfo(errorInfo ErrorInfo) {
    v.ErrorInfo = errorInfo
}

func (v *VoiceIdBatchFraudsterRegistrationAction) SetAction(action string) {
    v.Action = action
}

func (v *VoiceIdBatchFraudsterRegistrationAction) SetBatchJobId(batchJobId string) {
    v.BatchJobId = batchJobId
}

func (v *VoiceIdBatchFraudsterRegistrationAction) SetDomainId(domainId string) {
    v.DomainId = domainId
}

func (v *VoiceIdBatchFraudsterRegistrationAction) SetSourceId(sourceId string) {
    v.SourceId = sourceId
}

func (v *VoiceIdBatchFraudsterRegistrationAction) SetStatus(status string) {
    v.Status = status
}
