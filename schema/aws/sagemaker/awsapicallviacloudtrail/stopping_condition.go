package awsapicallviacloudtrail

type StoppingCondition struct {
    MaxRuntimeInSeconds float64 `json:"maxRuntimeInSeconds,omitempty"`
}

func (s *StoppingCondition) SetMaxRuntimeInSeconds(maxRuntimeInSeconds float64) {
    s.MaxRuntimeInSeconds = maxRuntimeInSeconds
}
