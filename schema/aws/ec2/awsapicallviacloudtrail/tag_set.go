package awsapicallviacloudtrail

type TagSet struct {
    Items []TagSpecificationSetItemItem `json:"items,omitempty"`
}

func (t *TagSet) SetItems(items []TagSpecificationSetItemItem) {
    t.Items = items
}
