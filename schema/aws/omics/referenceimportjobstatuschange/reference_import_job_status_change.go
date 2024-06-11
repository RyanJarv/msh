package referenceimportjobstatuschange

type ReferenceImportJobStatusChange struct {
    OmicsVersion string `json:"omicsVersion,omitempty"`
    ReferenceStoreId string `json:"referenceStoreId,omitempty"`
    Status string `json:"status,omitempty"`
    Id string `json:"id,omitempty"`
    StatusMessage string `json:"statusMessage,omitempty"`
}

func (r *ReferenceImportJobStatusChange) SetOmicsVersion(omicsVersion string) {
    r.OmicsVersion = omicsVersion
}

func (r *ReferenceImportJobStatusChange) SetReferenceStoreId(referenceStoreId string) {
    r.ReferenceStoreId = referenceStoreId
}

func (r *ReferenceImportJobStatusChange) SetStatus(status string) {
    r.Status = status
}

func (r *ReferenceImportJobStatusChange) SetId(id string) {
    r.Id = id
}

func (r *ReferenceImportJobStatusChange) SetStatusMessage(statusMessage string) {
    r.StatusMessage = statusMessage
}
