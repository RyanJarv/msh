package awsapicallviacloudtrail

type RequestParameters struct {
    Period float64 `json:"period,omitempty"`
    Statistic string `json:"statistic,omitempty"`
    MetricName string `json:"metricName,omitempty"`
    ComparisonOperator string `json:"comparisonOperator,omitempty"`
    Threshold float64 `json:"threshold,omitempty"`
    AlarmName string `json:"alarmName,omitempty"`
    EvaluationPeriods float64 `json:"evaluationPeriods,omitempty"`
    AlarmDescription string `json:"alarmDescription,omitempty"`
    Unit string `json:"unit,omitempty"`
    AlarmNames []string `json:"alarmNames,omitempty"`
    ActionsEnabled bool `json:"actionsEnabled,omitempty"`
    Namespace string `json:"namespace,omitempty"`
    AlarmActions []string `json:"alarmActions,omitempty"`
    Dimensions []RequestParametersItem `json:"dimensions,omitempty"`
}

func (r *RequestParameters) SetPeriod(period float64) {
    r.Period = period
}

func (r *RequestParameters) SetStatistic(statistic string) {
    r.Statistic = statistic
}

func (r *RequestParameters) SetMetricName(metricName string) {
    r.MetricName = metricName
}

func (r *RequestParameters) SetComparisonOperator(comparisonOperator string) {
    r.ComparisonOperator = comparisonOperator
}

func (r *RequestParameters) SetThreshold(threshold float64) {
    r.Threshold = threshold
}

func (r *RequestParameters) SetAlarmName(alarmName string) {
    r.AlarmName = alarmName
}

func (r *RequestParameters) SetEvaluationPeriods(evaluationPeriods float64) {
    r.EvaluationPeriods = evaluationPeriods
}

func (r *RequestParameters) SetAlarmDescription(alarmDescription string) {
    r.AlarmDescription = alarmDescription
}

func (r *RequestParameters) SetUnit(unit string) {
    r.Unit = unit
}

func (r *RequestParameters) SetAlarmNames(alarmNames []string) {
    r.AlarmNames = alarmNames
}

func (r *RequestParameters) SetActionsEnabled(actionsEnabled bool) {
    r.ActionsEnabled = actionsEnabled
}

func (r *RequestParameters) SetNamespace(namespace string) {
    r.Namespace = namespace
}

func (r *RequestParameters) SetAlarmActions(alarmActions []string) {
    r.AlarmActions = alarmActions
}

func (r *RequestParameters) SetDimensions(dimensions []RequestParametersItem) {
    r.Dimensions = dimensions
}
