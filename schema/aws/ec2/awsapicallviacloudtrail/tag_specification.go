package awsapicallviacloudtrail

type TagSpecification struct {
    Tag Tag `json:"Tag,omitempty"`
    ResourceType string `json:"ResourceType,omitempty"`
}

func (t *TagSpecification) SetTag(tag Tag) {
    t.Tag = tag
}

func (t *TagSpecification) SetResourceType(resourceType string) {
    t.ResourceType = resourceType
}
