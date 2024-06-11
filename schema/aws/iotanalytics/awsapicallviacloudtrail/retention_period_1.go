package awsapicallviacloudtrail

type RetentionPeriod_1 struct {
    NumberOfDays float64 `json:"numberOfDays,omitempty"`
    Unlimited bool `json:"unlimited,omitempty"`
}

func (r *RetentionPeriod_1) SetNumberOfDays(numberOfDays float64) {
    r.NumberOfDays = numberOfDays
}

func (r *RetentionPeriod_1) SetUnlimited(unlimited bool) {
    r.Unlimited = unlimited
}
