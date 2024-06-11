package sagemakerhyperparametertuningjobstatechange

type StoppingCondition struct {
    MaxRuntimeInSeconds float64 `json:"MaxRuntimeInSeconds"`
}

func (s *StoppingCondition) SetMaxRuntimeInSeconds(maxRuntimeInSeconds float64) {
    s.MaxRuntimeInSeconds = maxRuntimeInSeconds
}
