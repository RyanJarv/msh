package readsetexportjobstatuschange

type ReadSetExportJobStatusChange struct {
    OmicsVersion string `json:"omicsVersion,omitempty"`
    SequenceStoreId string `json:"sequenceStoreId,omitempty"`
    Status string `json:"status,omitempty"`
    Id string `json:"id,omitempty"`
    StatusMessage string `json:"statusMessage,omitempty"`
}

func (r *ReadSetExportJobStatusChange) SetOmicsVersion(omicsVersion string) {
    r.OmicsVersion = omicsVersion
}

func (r *ReadSetExportJobStatusChange) SetSequenceStoreId(sequenceStoreId string) {
    r.SequenceStoreId = sequenceStoreId
}

func (r *ReadSetExportJobStatusChange) SetStatus(status string) {
    r.Status = status
}

func (r *ReadSetExportJobStatusChange) SetId(id string) {
    r.Id = id
}

func (r *ReadSetExportJobStatusChange) SetStatusMessage(statusMessage string) {
    r.StatusMessage = statusMessage
}
