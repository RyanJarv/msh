package runstatuschange

type RunStatusChange struct {
    OmicsVersion string `json:"omicsVersion,omitempty"`
    Arn string `json:"arn,omitempty"`
    Status string `json:"status,omitempty"`
}

func (r *RunStatusChange) SetOmicsVersion(omicsVersion string) {
    r.OmicsVersion = omicsVersion
}

func (r *RunStatusChange) SetArn(arn string) {
    r.Arn = arn
}

func (r *RunStatusChange) SetStatus(status string) {
    r.Status = status
}
