package readsetactivationjobstatuschange

type ReadSetActivationJobStatusChange struct {
    OmicsVersion string `json:"omicsVersion,omitempty"`
    SequenceStoreId string `json:"sequenceStoreId,omitempty"`
    Status string `json:"status,omitempty"`
    Id string `json:"id,omitempty"`
    StatusMessage string `json:"statusMessage,omitempty"`
}

func (r *ReadSetActivationJobStatusChange) SetOmicsVersion(omicsVersion string) {
    r.OmicsVersion = omicsVersion
}

func (r *ReadSetActivationJobStatusChange) SetSequenceStoreId(sequenceStoreId string) {
    r.SequenceStoreId = sequenceStoreId
}

func (r *ReadSetActivationJobStatusChange) SetStatus(status string) {
    r.Status = status
}

func (r *ReadSetActivationJobStatusChange) SetId(id string) {
    r.Id = id
}

func (r *ReadSetActivationJobStatusChange) SetStatusMessage(statusMessage string) {
    r.StatusMessage = statusMessage
}
