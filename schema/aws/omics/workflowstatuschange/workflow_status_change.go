package workflowstatuschange

type WorkflowStatusChange struct {
    OmicsVersion string `json:"omicsVersion,omitempty"`
    Arn string `json:"arn,omitempty"`
    Status string `json:"status,omitempty"`
}

func (w *WorkflowStatusChange) SetOmicsVersion(omicsVersion string) {
    w.OmicsVersion = omicsVersion
}

func (w *WorkflowStatusChange) SetArn(arn string) {
    w.Arn = arn
}

func (w *WorkflowStatusChange) SetStatus(status string) {
    w.Status = status
}
