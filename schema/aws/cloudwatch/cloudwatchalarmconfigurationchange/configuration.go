package cloudwatchalarmconfigurationchange

type Configuration struct {
    AlarmName string `json:"alarmName"`
    Description string `json:"description,omitempty"`
    ActionsEnabled bool `json:"actionsEnabled"`
    Timestamp string `json:"timestamp"`
    EvaluationPeriods float64 `json:"evaluationPeriods,omitempty"`
    DatapointsToAlarm float64 `json:"datapointsToAlarm,omitempty"`
    Threshold float64 `json:"threshold,omitempty"`
    ComparisonOperator string `json:"comparisonOperator,omitempty"`
    TreatMissingData string `json:"treatMissingData,omitempty"`
    EvaluateLowSampleCountPercentile string `json:"evaluateLowSampleCountPercentile,omitempty"`
    Metrics []ConfigurationItem `json:"metrics,omitempty"`
    AlarmRule string `json:"alarmRule,omitempty"`
    OkActions []string `json:"okActions"`
    AlarmActions []string `json:"alarmActions"`
    InsufficientDataActions []string `json:"insufficientDataActions"`
    ActionsSuppressor string `json:"actionsSuppressor,omitempty"`
    ActionsSuppressorWaitPeriod float64 `json:"actionsSuppressorWaitPeriod,omitempty"`
    ActionsSuppressorExtensionPeriod float64 `json:"actionsSuppressorExtensionPeriod,omitempty"`
}

func (c *Configuration) SetAlarmName(alarmName string) {
    c.AlarmName = alarmName
}

func (c *Configuration) SetDescription(description string) {
    c.Description = description
}

func (c *Configuration) SetActionsEnabled(actionsEnabled bool) {
    c.ActionsEnabled = actionsEnabled
}

func (c *Configuration) SetTimestamp(timestamp string) {
    c.Timestamp = timestamp
}

func (c *Configuration) SetEvaluationPeriods(evaluationPeriods float64) {
    c.EvaluationPeriods = evaluationPeriods
}

func (c *Configuration) SetDatapointsToAlarm(datapointsToAlarm float64) {
    c.DatapointsToAlarm = datapointsToAlarm
}

func (c *Configuration) SetThreshold(threshold float64) {
    c.Threshold = threshold
}

func (c *Configuration) SetComparisonOperator(comparisonOperator string) {
    c.ComparisonOperator = comparisonOperator
}

func (c *Configuration) SetTreatMissingData(treatMissingData string) {
    c.TreatMissingData = treatMissingData
}

func (c *Configuration) SetEvaluateLowSampleCountPercentile(evaluateLowSampleCountPercentile string) {
    c.EvaluateLowSampleCountPercentile = evaluateLowSampleCountPercentile
}

func (c *Configuration) SetMetrics(metrics []ConfigurationItem) {
    c.Metrics = metrics
}

func (c *Configuration) SetAlarmRule(alarmRule string) {
    c.AlarmRule = alarmRule
}

func (c *Configuration) SetOkActions(okActions []string) {
    c.OkActions = okActions
}

func (c *Configuration) SetAlarmActions(alarmActions []string) {
    c.AlarmActions = alarmActions
}

func (c *Configuration) SetInsufficientDataActions(insufficientDataActions []string) {
    c.InsufficientDataActions = insufficientDataActions
}

func (c *Configuration) SetActionsSuppressor(actionsSuppressor string) {
    c.ActionsSuppressor = actionsSuppressor
}

func (c *Configuration) SetActionsSuppressorWaitPeriod(actionsSuppressorWaitPeriod float64) {
    c.ActionsSuppressorWaitPeriod = actionsSuppressorWaitPeriod
}

func (c *Configuration) SetActionsSuppressorExtensionPeriod(actionsSuppressorExtensionPeriod float64) {
    c.ActionsSuppressorExtensionPeriod = actionsSuppressorExtensionPeriod
}
