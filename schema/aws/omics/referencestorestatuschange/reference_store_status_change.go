package referencestorestatuschange

type ReferenceStoreStatusChange struct {
    OmicsVersion string `json:"omicsVersion,omitempty"`
    Arn string `json:"arn,omitempty"`
    Status string `json:"status,omitempty"`
    Id string `json:"id,omitempty"`
}

func (r *ReferenceStoreStatusChange) SetOmicsVersion(omicsVersion string) {
    r.OmicsVersion = omicsVersion
}

func (r *ReferenceStoreStatusChange) SetArn(arn string) {
    r.Arn = arn
}

func (r *ReferenceStoreStatusChange) SetStatus(status string) {
    r.Status = status
}

func (r *ReferenceStoreStatusChange) SetId(id string) {
    r.Id = id
}
