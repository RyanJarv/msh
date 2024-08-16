package cloudwatchalarmstatechange

type State struct {
    Reason string `json:"reason"`
    ReasonData string `json:"reasonData"`
    Value string `json:"value"`
    Timestamp string `json:"timestamp"`
    ActionsSuppressedBy string `json:"actionsSuppressedBy,omitempty"`
    ActionsSuppressedReason string `json:"actionsSuppressedReason,omitempty"`
    EvaluationState string `json:"evaluationState,omitempty"`
}

func (s *State) SetReason(reason string) {
    s.Reason = reason
}

func (s *State) SetReasonData(reasonData string) {
    s.ReasonData = reasonData
}

func (s *State) SetValue(value string) {
    s.Value = value
}

func (s *State) SetTimestamp(timestamp string) {
    s.Timestamp = timestamp
}

func (s *State) SetActionsSuppressedBy(actionsSuppressedBy string) {
    s.ActionsSuppressedBy = actionsSuppressedBy
}

func (s *State) SetActionsSuppressedReason(actionsSuppressedReason string) {
    s.ActionsSuppressedReason = actionsSuppressedReason
}

func (s *State) SetEvaluationState(evaluationState string) {
    s.EvaluationState = evaluationState
}