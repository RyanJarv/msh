package awsapicallviacloudtrail

type CreateRule struct {
    Interval float64 `json:"Interval"`
    IntervalUnit string `json:"IntervalUnit"`
    Times []string `json:"Times"`
}

func (c *CreateRule) SetInterval(interval float64) {
    c.Interval = interval
}

func (c *CreateRule) SetIntervalUnit(intervalUnit string) {
    c.IntervalUnit = intervalUnit
}

func (c *CreateRule) SetTimes(times []string) {
    c.Times = times
}
