package readsetimportjobstatuschange

type ReadSetImportJobStatusChange struct {
    OmicsVersion string `json:"omicsVersion,omitempty"`
    SequenceStoreId string `json:"sequenceStoreId,omitempty"`
    Status string `json:"status,omitempty"`
    Id string `json:"id,omitempty"`
    StatusMessage string `json:"statusMessage,omitempty"`
}

func (r *ReadSetImportJobStatusChange) SetOmicsVersion(omicsVersion string) {
    r.OmicsVersion = omicsVersion
}

func (r *ReadSetImportJobStatusChange) SetSequenceStoreId(sequenceStoreId string) {
    r.SequenceStoreId = sequenceStoreId
}

func (r *ReadSetImportJobStatusChange) SetStatus(status string) {
    r.Status = status
}

func (r *ReadSetImportJobStatusChange) SetId(id string) {
    r.Id = id
}

func (r *ReadSetImportJobStatusChange) SetStatusMessage(statusMessage string) {
    r.StatusMessage = statusMessage
}
