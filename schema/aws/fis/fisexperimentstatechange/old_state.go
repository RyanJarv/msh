package fisexperimentstatechange

type Old_state struct {
    Reason string `json:"reason"`
    Status string `json:"status"`
}

func (o *Old_state) SetReason(reason string) {
    o.Reason = reason
}

func (o *Old_state) SetStatus(status string) {
    o.Status = status
}
