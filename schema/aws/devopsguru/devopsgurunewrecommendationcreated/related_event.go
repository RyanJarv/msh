package devopsgurunewrecommendationcreated

type RelatedEvent struct {
    Name string `json:"name"`
    Resources []EventResource `json:"resources"`
}

func (r *RelatedEvent) SetName(name string) {
    r.Name = name
}

func (r *RelatedEvent) SetResources(resources []EventResource) {
    r.Resources = resources
}
