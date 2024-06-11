package cloudwatchalarmstatechange

type Configuration struct {
    Metrics []ConfigurationItem `json:"metrics,omitempty"`
    Description string `json:"description,omitempty"`
    AlarmRule string `json:"alarmRule,omitempty"`
    ActionsSuppressor string `json:"actionsSuppressor,omitempty"`
    ActionsSuppressorWaitPeriod float64 `json:"actionsSuppressorWaitPeriod,omitempty"`
    ActionsSuppressorExtensionPeriod float64 `json:"actionsSuppressorExtensionPeriod,omitempty"`
}

func (c *Configuration) SetMetrics(metrics []ConfigurationItem) {
    c.Metrics = metrics
}

func (c *Configuration) SetDescription(description string) {
    c.Description = description
}

func (c *Configuration) SetAlarmRule(alarmRule string) {
    c.AlarmRule = alarmRule
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
