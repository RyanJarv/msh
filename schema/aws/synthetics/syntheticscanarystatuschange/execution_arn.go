package syntheticscanarystatuschange

type ExecutionArn struct {
    CurrentValue string `json:"current-value"`
    PreviousValue string `json:"previous-value"`
}

func (e *ExecutionArn) SetCurrentValue(currentValue string) {
    e.CurrentValue = currentValue
}

func (e *ExecutionArn) SetPreviousValue(previousValue string) {
    e.PreviousValue = previousValue
}
