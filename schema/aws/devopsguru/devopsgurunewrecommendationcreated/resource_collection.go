package devopsgurunewrecommendationcreated

type ResourceCollection struct {
    CloudFormation CloudFormation `json:"cloudFormation,omitempty"`
    Tags []Tag `json:"tags,omitempty"`
}

func (r *ResourceCollection) SetCloudFormation(cloudFormation CloudFormation) {
    r.CloudFormation = cloudFormation
}

func (r *ResourceCollection) SetTags(tags []Tag) {
    r.Tags = tags
}
