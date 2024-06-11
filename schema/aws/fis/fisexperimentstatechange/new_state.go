package fisexperimentstatechange

type New_state struct {
    Reason string `json:"reason"`
    Status string `json:"status"`
}

func (n *New_state) SetReason(reason string) {
    n.Reason = reason
}

func (n *New_state) SetStatus(status string) {
    n.Status = status
}
