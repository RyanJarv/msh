package referencestatuschange

type ReferenceStatusChange struct {
    OmicsVersion string `json:"omicsVersion,omitempty"`
    ReferenceStoreId string `json:"referenceStoreId,omitempty"`
    Status string `json:"status,omitempty"`
    Id string `json:"id,omitempty"`
    Arn string `json:"arn,omitempty"`
}

func (r *ReferenceStatusChange) SetOmicsVersion(omicsVersion string) {
    r.OmicsVersion = omicsVersion
}

func (r *ReferenceStatusChange) SetReferenceStoreId(referenceStoreId string) {
    r.ReferenceStoreId = referenceStoreId
}

func (r *ReferenceStatusChange) SetStatus(status string) {
    r.Status = status
}

func (r *ReferenceStatusChange) SetId(id string) {
    r.Id = id
}

func (r *ReferenceStatusChange) SetArn(arn string) {
    r.Arn = arn
}
