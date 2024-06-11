package route53applicationrecoverycontrollerrecoverygroupreadinessstatuschange

type RecoveryGroupStatusChange struct {
    NewState State `json:"new-state"`
    PreviousState State `json:"previous-state"`
    RecoveryGroupName string `json:"recovery-group-name"`
}

func (r *RecoveryGroupStatusChange) SetNewState(newState State) {
    r.NewState = newState
}

func (r *RecoveryGroupStatusChange) SetPreviousState(previousState State) {
    r.PreviousState = previousState
}

func (r *RecoveryGroupStatusChange) SetRecoveryGroupName(recoveryGroupName string) {
    r.RecoveryGroupName = recoveryGroupName
}
