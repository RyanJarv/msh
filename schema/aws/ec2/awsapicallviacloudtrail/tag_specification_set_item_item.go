package awsapicallviacloudtrail

type TagSpecificationSetItemItem struct {
    Value string `json:"value,omitempty"`
    Key string `json:"key,omitempty"`
}

func (t *TagSpecificationSetItemItem) SetValue(value string) {
    t.Value = value
}

func (t *TagSpecificationSetItemItem) SetKey(key string) {
    t.Key = key
}
