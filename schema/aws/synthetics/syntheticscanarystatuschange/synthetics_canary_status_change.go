package syntheticscanarystatuschange

type SyntheticsCanaryStatusChange struct {
    ChangedConfig Changed_config `json:"changed-config"`
    AccountId string `json:"account-id"`
    CanaryId string `json:"canary-id"`
    CanaryName string `json:"canary-name"`
    CurrentState string `json:"current-state"`
    Message string `json:"message"`
    PreviousState string `json:"previous-state"`
    SourceLocation string `json:"source-location"`
    UpdatedOn float64 `json:"updated-on"`
}

func (s *SyntheticsCanaryStatusChange) SetChangedConfig(changedConfig Changed_config) {
    s.ChangedConfig = changedConfig
}

func (s *SyntheticsCanaryStatusChange) SetAccountId(accountId string) {
    s.AccountId = accountId
}

func (s *SyntheticsCanaryStatusChange) SetCanaryId(canaryId string) {
    s.CanaryId = canaryId
}

func (s *SyntheticsCanaryStatusChange) SetCanaryName(canaryName string) {
    s.CanaryName = canaryName
}

func (s *SyntheticsCanaryStatusChange) SetCurrentState(currentState string) {
    s.CurrentState = currentState
}

func (s *SyntheticsCanaryStatusChange) SetMessage(message string) {
    s.Message = message
}

func (s *SyntheticsCanaryStatusChange) SetPreviousState(previousState string) {
    s.PreviousState = previousState
}

func (s *SyntheticsCanaryStatusChange) SetSourceLocation(sourceLocation string) {
    s.SourceLocation = sourceLocation
}

func (s *SyntheticsCanaryStatusChange) SetUpdatedOn(updatedOn float64) {
    s.UpdatedOn = updatedOn
}
