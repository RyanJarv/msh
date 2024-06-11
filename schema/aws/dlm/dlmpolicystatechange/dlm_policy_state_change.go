package dlmpolicystatechange

type DLMPolicyStateChange struct {
    PolicyId string `json:"policy_id"`
    Cause string `json:"cause"`
    State string `json:"state"`
}

func (d *DLMPolicyStateChange) SetPolicyId(policyId string) {
    d.PolicyId = policyId
}

func (d *DLMPolicyStateChange) SetCause(cause string) {
    d.Cause = cause
}

func (d *DLMPolicyStateChange) SetState(state string) {
    d.State = state
}
