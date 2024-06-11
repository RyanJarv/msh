package route53applicationrecoverycontrollerreadinesscheckstatuschange

type ReadinessCheckStatusChange struct {
    NewState State `json:"new-state"`
    PreviousState State `json:"previous-state"`
    ReadinessCheckName string `json:"readiness-check-name"`
}

func (r *ReadinessCheckStatusChange) SetNewState(newState State) {
    r.NewState = newState
}

func (r *ReadinessCheckStatusChange) SetPreviousState(previousState State) {
    r.PreviousState = previousState
}

func (r *ReadinessCheckStatusChange) SetReadinessCheckName(readinessCheckName string) {
    r.ReadinessCheckName = readinessCheckName
}
