package awsapicallviacloudtrail

type PolicyDetailsItem struct {
    CreateRule CreateRule `json:"CreateRule"`
    FastRestoreRule FastRestoreRule `json:"FastRestoreRule,omitempty"`
    RetainRule RetainRule `json:"RetainRule"`
    CopyTags bool `json:"CopyTags,omitempty"`
    Name string `json:"Name"`
    TagsToAdd []PolicyDetailsItemItem `json:"TagsToAdd"`
}

func (p *PolicyDetailsItem) SetCreateRule(createRule CreateRule) {
    p.CreateRule = createRule
}

func (p *PolicyDetailsItem) SetFastRestoreRule(fastRestoreRule FastRestoreRule) {
    p.FastRestoreRule = fastRestoreRule
}

func (p *PolicyDetailsItem) SetRetainRule(retainRule RetainRule) {
    p.RetainRule = retainRule
}

func (p *PolicyDetailsItem) SetCopyTags(copyTags bool) {
    p.CopyTags = copyTags
}

func (p *PolicyDetailsItem) SetName(name string) {
    p.Name = name
}

func (p *PolicyDetailsItem) SetTagsToAdd(tagsToAdd []PolicyDetailsItemItem) {
    p.TagsToAdd = tagsToAdd
}
