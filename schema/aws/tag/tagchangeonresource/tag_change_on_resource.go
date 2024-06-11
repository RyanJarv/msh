package tagchangeonresource

type TagChangeOnResource struct {
    Tags Tags `json:"tags"`
    ChangedTagKeys []string `json:"changed-tag-keys"`
    Service string `json:"service"`
    ResourceType string `json:"resource-type"`
    Version float64 `json:"version"`
    TagPolicyCompliant string `json:"tag-policy-compliant,omitempty"`
    VersionTimestamp string `json:"version-timestamp,omitempty"`
}

func (t *TagChangeOnResource) SetTags(tags Tags) {
    t.Tags = tags
}

func (t *TagChangeOnResource) SetChangedTagKeys(changedTagKeys []string) {
    t.ChangedTagKeys = changedTagKeys
}

func (t *TagChangeOnResource) SetService(service string) {
    t.Service = service
}

func (t *TagChangeOnResource) SetResourceType(resourceType string) {
    t.ResourceType = resourceType
}

func (t *TagChangeOnResource) SetVersion(version float64) {
    t.Version = version
}

func (t *TagChangeOnResource) SetTagPolicyCompliant(tagPolicyCompliant string) {
    t.TagPolicyCompliant = tagPolicyCompliant
}

func (t *TagChangeOnResource) SetVersionTimestamp(versionTimestamp string) {
    t.VersionTimestamp = versionTimestamp
}
