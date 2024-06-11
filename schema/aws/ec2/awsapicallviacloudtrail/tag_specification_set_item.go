package awsapicallviacloudtrail

type TagSpecificationSetItem struct {
    ResourceType string `json:"resourceType"`
    Tags []TagSpecificationSetItemItem `json:"tags"`
}

func (t *TagSpecificationSetItem) SetResourceType(resourceType string) {
    t.ResourceType = resourceType
}

func (t *TagSpecificationSetItem) SetTags(tags []TagSpecificationSetItemItem) {
    t.Tags = tags
}
