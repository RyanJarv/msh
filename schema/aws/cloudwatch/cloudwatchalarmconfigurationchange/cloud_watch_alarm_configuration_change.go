package cloudwatchalarmconfigurationchange

type CloudWatchAlarmConfigurationChange struct {
    AlarmName string `json:"alarmName"`
    Operation string `json:"operation"`
    Configuration Configuration `json:"configuration"`
    State State `json:"state"`
}

func (c *CloudWatchAlarmConfigurationChange) SetAlarmName(alarmName string) {
    c.AlarmName = alarmName
}

func (c *CloudWatchAlarmConfigurationChange) SetOperation(operation string) {
    c.Operation = operation
}

func (c *CloudWatchAlarmConfigurationChange) SetConfiguration(configuration Configuration) {
    c.Configuration = configuration
}

func (c *CloudWatchAlarmConfigurationChange) SetState(state State) {
    c.State = state
}
