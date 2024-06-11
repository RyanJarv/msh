package awsapicallviacloudtrail

type RetentionPeriod struct {
    NumberOfDays float64 `json:"numberOfDays,omitempty"`
    Unlimited bool `json:"unlimited,omitempty"`
}

func (r *RetentionPeriod) SetNumberOfDays(numberOfDays float64) {
    r.NumberOfDays = numberOfDays
}

func (r *RetentionPeriod) SetUnlimited(unlimited bool) {
    r.Unlimited = unlimited
}
