package devopsguruinsightclosed

type ReferenceValue struct {
    Threshold float64 `json:"threshold,omitempty"`
    ComparisonOperator string `json:"comparisonOperator,omitempty"`
    DatapointsToAlarm float64 `json:"datapointsToAlarm,omitempty"`
    EvaluationPeriod float64 `json:"evaluationPeriod,omitempty"`
}

func (r *ReferenceValue) SetThreshold(threshold float64) {
    r.Threshold = threshold
}

func (r *ReferenceValue) SetComparisonOperator(comparisonOperator string) {
    r.ComparisonOperator = comparisonOperator
}

func (r *ReferenceValue) SetDatapointsToAlarm(datapointsToAlarm float64) {
    r.DatapointsToAlarm = datapointsToAlarm
}

func (r *ReferenceValue) SetEvaluationPeriod(evaluationPeriod float64) {
    r.EvaluationPeriod = evaluationPeriod
}
