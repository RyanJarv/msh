package devopsgurunewrecommendationcreated

type Recommendation struct {
    Description string `json:"description"`
    Link string `json:"link"`
    Name string `json:"name"`
    Reason string `json:"reason"`
    RelatedAnomalies []RelatedAnomaly `json:"relatedAnomalies"`
    RelatedEvents []RelatedEvent `json:"relatedEvents"`
}

func (r *Recommendation) SetDescription(description string) {
    r.Description = description
}

func (r *Recommendation) SetLink(link string) {
    r.Link = link
}

func (r *Recommendation) SetName(name string) {
    r.Name = name
}

func (r *Recommendation) SetReason(reason string) {
    r.Reason = reason
}

func (r *Recommendation) SetRelatedAnomalies(relatedAnomalies []RelatedAnomaly) {
    r.RelatedAnomalies = relatedAnomalies
}

func (r *Recommendation) SetRelatedEvents(relatedEvents []RelatedEvent) {
    r.RelatedEvents = relatedEvents
}
