package devopsgurunewanomalyassociation

type Tag struct {
    AppBoundaryKey string `json:"appBoundaryKey,omitempty"`
    TagValues []string `json:"tagValues,omitempty"`
}

func (t *Tag) SetAppBoundaryKey(appBoundaryKey string) {
    t.AppBoundaryKey = appBoundaryKey
}

func (t *Tag) SetTagValues(tagValues []string) {
    t.TagValues = tagValues
}
