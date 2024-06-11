package cloudwatchalarmstatechange

type CloudWatchAlarmStateChange struct {
    Configuration Configuration `json:"configuration"`
    State State `json:"state"`
    PreviousState State `json:"previousState"`
    AlarmName string `json:"alarmName"`
}

func (c *CloudWatchAlarmStateChange) SetConfiguration(configuration Configuration) {
    c.Configuration = configuration
}

func (c *CloudWatchAlarmStateChange) SetState(state State) {
    c.State = state
}

func (c *CloudWatchAlarmStateChange) SetPreviousState(previousState State) {
    c.PreviousState = previousState
}

func (c *CloudWatchAlarmStateChange) SetAlarmName(alarmName string) {
    c.AlarmName = alarmName
}
