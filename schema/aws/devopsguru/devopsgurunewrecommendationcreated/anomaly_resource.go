package devopsgurunewrecommendationcreated

type AnomalyResource struct {
    Name string `json:"name"`
    AnomalyResourceType string `json:"type"`
}

func (a *AnomalyResource) SetName(name string) {
    a.Name = name
}

func (a *AnomalyResource) SetAnomalyResourceType(anomalyResourceType string) {
    a.AnomalyResourceType = anomalyResourceType
}
