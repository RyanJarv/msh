package syntheticscanarystatuschange

type TestCodeLayerVersionArn struct {
    CurrentValue string `json:"current-value"`
    PreviousValue string `json:"previous-value"`
}

func (t *TestCodeLayerVersionArn) SetCurrentValue(currentValue string) {
    t.CurrentValue = currentValue
}

func (t *TestCodeLayerVersionArn) SetPreviousValue(previousValue string) {
    t.PreviousValue = previousValue
}
