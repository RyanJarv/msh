package awshealthabuseevent

type AffectedEntity struct {
    Tags Tags `json:"tags,omitempty"`
    EntityValue string `json:"entityValue,omitempty"`
}

func (a *AffectedEntity) SetTags(tags Tags) {
    a.Tags = tags
}

func (a *AffectedEntity) SetEntityValue(entityValue string) {
    a.EntityValue = entityValue
}
