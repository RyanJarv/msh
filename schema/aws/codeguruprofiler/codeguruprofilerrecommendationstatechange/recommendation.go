package codeguruprofilerrecommendationstatechange

type Recommendation struct {
    Description Description `json:"description"`
    Name Name `json:"name"`
    Reason Reason `json:"reason"`
    ResolutionSteps ResolutionSteps `json:"resolutionSteps"`
}

func (r *Recommendation) SetDescription(description Description) {
    r.Description = description
}

func (r *Recommendation) SetName(name Name) {
    r.Name = name
}

func (r *Recommendation) SetReason(reason Reason) {
    r.Reason = reason
}

func (r *Recommendation) SetResolutionSteps(resolutionSteps ResolutionSteps) {
    r.ResolutionSteps = resolutionSteps
}
