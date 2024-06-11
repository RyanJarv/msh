package readsetstatuschange

type ReadSetStatusChange struct {
    OmicsVersion string `json:"omicsVersion,omitempty"`
    SequenceStoreId string `json:"sequenceStoreId,omitempty"`
    Status string `json:"status,omitempty"`
    Id string `json:"id,omitempty"`
    StatusMessage string `json:"statusMessage,omitempty"`
    Arn string `json:"arn,omitempty"`
}

func (r *ReadSetStatusChange) SetOmicsVersion(omicsVersion string) {
    r.OmicsVersion = omicsVersion
}

func (r *ReadSetStatusChange) SetSequenceStoreId(sequenceStoreId string) {
    r.SequenceStoreId = sequenceStoreId
}

func (r *ReadSetStatusChange) SetStatus(status string) {
    r.Status = status
}

func (r *ReadSetStatusChange) SetId(id string) {
    r.Id = id
}

func (r *ReadSetStatusChange) SetStatusMessage(statusMessage string) {
    r.StatusMessage = statusMessage
}

func (r *ReadSetStatusChange) SetArn(arn string) {
    r.Arn = arn
}
