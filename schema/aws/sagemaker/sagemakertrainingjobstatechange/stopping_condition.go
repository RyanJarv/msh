package sagemakertrainingjobstatechange

type StoppingCondition struct {
    MaxRuntimeInSeconds float64 `json:"MaxRuntimeInSeconds,omitempty"`
}

func (s *StoppingCondition) SetMaxRuntimeInSeconds(maxRuntimeInSeconds float64) {
    s.MaxRuntimeInSeconds = maxRuntimeInSeconds
}
