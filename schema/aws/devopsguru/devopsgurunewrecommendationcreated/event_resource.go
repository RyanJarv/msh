package devopsgurunewrecommendationcreated

type EventResource struct {
    Name string `json:"name"`
    EventResourceType string `json:"type"`
}

func (e *EventResource) SetName(name string) {
    e.Name = name
}

func (e *EventResource) SetEventResourceType(eventResourceType string) {
    e.EventResourceType = eventResourceType
}
