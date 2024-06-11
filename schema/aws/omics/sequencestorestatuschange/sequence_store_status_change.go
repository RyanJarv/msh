package sequencestorestatuschange

type SequenceStoreStatusChange struct {
    OmicsVersion string `json:"omicsVersion,omitempty"`
    Arn string `json:"arn,omitempty"`
    Status string `json:"status,omitempty"`
    Id string `json:"id,omitempty"`
}

func (s *SequenceStoreStatusChange) SetOmicsVersion(omicsVersion string) {
    s.OmicsVersion = omicsVersion
}

func (s *SequenceStoreStatusChange) SetArn(arn string) {
    s.Arn = arn
}

func (s *SequenceStoreStatusChange) SetStatus(status string) {
    s.Status = status
}

func (s *SequenceStoreStatusChange) SetId(id string) {
    s.Id = id
}
