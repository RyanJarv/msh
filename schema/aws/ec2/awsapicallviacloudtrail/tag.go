package awsapicallviacloudtrail

type Tag struct {
    Value string `json:"Value,omitempty"`
    Tag float64 `json:"tag,omitempty"`
    Key string `json:"Key,omitempty"`
}

func (t *Tag) SetValue(value string) {
    t.Value = value
}

func (t *Tag) SetTag(tag float64) {
    t.Tag = tag
}

func (t *Tag) SetKey(key string) {
    t.Key = key
}
