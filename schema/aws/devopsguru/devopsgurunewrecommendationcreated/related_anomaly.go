package devopsgurunewrecommendationcreated

type RelatedAnomaly struct {
    SourceDetails []RelatedAnomalySourceDetail `json:"sourceDetails"`
    Resources []AnomalyResource `json:"resources"`
    AssociatedResourceArns []string `json:"associatedResourceArns,omitempty"`
}

func (r *RelatedAnomaly) SetSourceDetails(sourceDetails []RelatedAnomalySourceDetail) {
    r.SourceDetails = sourceDetails
}

func (r *RelatedAnomaly) SetResources(resources []AnomalyResource) {
    r.Resources = resources
}

func (r *RelatedAnomaly) SetAssociatedResourceArns(associatedResourceArns []string) {
    r.AssociatedResourceArns = associatedResourceArns
}
