package emrclusterstatechange

type EMRClusterStateChange struct {
    Severity string `json:"severity,omitempty"`
    StateChangeReason string `json:"stateChangeReason,omitempty"`
    Name string `json:"name,omitempty"`
    ClusterId string `json:"clusterId,omitempty"`
    State string `json:"state,omitempty"`
    Message string `json:"message,omitempty"`
}

func (e *EMRClusterStateChange) SetSeverity(severity string) {
    e.Severity = severity
}

func (e *EMRClusterStateChange) SetStateChangeReason(stateChangeReason string) {
    e.StateChangeReason = stateChangeReason
}

func (e *EMRClusterStateChange) SetName(name string) {
    e.Name = name
}

func (e *EMRClusterStateChange) SetClusterId(clusterId string) {
    e.ClusterId = clusterId
}

func (e *EMRClusterStateChange) SetState(state string) {
    e.State = state
}

func (e *EMRClusterStateChange) SetMessage(message string) {
    e.Message = message
}
