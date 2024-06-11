package taskstatuschange

type TaskStatusChange struct {
    OmicsVersion string `json:"omicsVersion,omitempty"`
    Arn string `json:"arn,omitempty"`
    Status string `json:"status,omitempty"`
}

func (t *TaskStatusChange) SetOmicsVersion(omicsVersion string) {
    t.OmicsVersion = omicsVersion
}

func (t *TaskStatusChange) SetArn(arn string) {
    t.Arn = arn
}

func (t *TaskStatusChange) SetStatus(status string) {
    t.Status = status
}
