package devopsguruinsightclosed

type AnomalyResource struct {
    Name string `json:"name,omitempty"`
    AnomalyResourceType string `json:"type,omitempty"`
}

func (a *AnomalyResource) SetName(name string) {
    a.Name = name
}

func (a *AnomalyResource) SetAnomalyResourceType(anomalyResourceType string) {
    a.AnomalyResourceType = anomalyResourceType
}
