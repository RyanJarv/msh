package codeguruprofilerrecommendationstatechange

type ResolutionSteps struct {
    Value string `json:"value"`
}

func (r *ResolutionSteps) SetValue(value string) {
    r.Value = value
}
