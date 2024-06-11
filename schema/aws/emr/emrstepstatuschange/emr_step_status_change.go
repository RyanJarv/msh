package emrstepstatuschange

type EMRStepStatusChange struct {
    Severity string `json:"severity,omitempty"`
    ActionOnFailure string `json:"actionOnFailure,omitempty"`
    StepId string `json:"stepId,omitempty"`
    Name string `json:"name,omitempty"`
    ClusterId string `json:"clusterId,omitempty"`
    State string `json:"state,omitempty"`
    Message string `json:"message,omitempty"`
}

func (e *EMRStepStatusChange) SetSeverity(severity string) {
    e.Severity = severity
}

func (e *EMRStepStatusChange) SetActionOnFailure(actionOnFailure string) {
    e.ActionOnFailure = actionOnFailure
}

func (e *EMRStepStatusChange) SetStepId(stepId string) {
    e.StepId = stepId
}

func (e *EMRStepStatusChange) SetName(name string) {
    e.Name = name
}

func (e *EMRStepStatusChange) SetClusterId(clusterId string) {
    e.ClusterId = clusterId
}

func (e *EMRStepStatusChange) SetState(state string) {
    e.State = state
}

func (e *EMRStepStatusChange) SetMessage(message string) {
    e.Message = message
}
