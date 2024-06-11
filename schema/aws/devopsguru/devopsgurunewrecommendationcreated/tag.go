package devopsgurunewrecommendationcreated

type Tag struct {
    AppBoundaryKey string `json:"appBoundaryKey"`
    TagValues []string `json:"tagValues"`
}

func (t *Tag) SetAppBoundaryKey(appBoundaryKey string) {
    t.AppBoundaryKey = appBoundaryKey
}

func (t *Tag) SetTagValues(tagValues []string) {
    t.TagValues = tagValues
}
