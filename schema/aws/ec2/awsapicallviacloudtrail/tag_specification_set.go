package awsapicallviacloudtrail

type TagSpecificationSet struct {
    Items []TagSpecificationSetItem `json:"items,omitempty"`
}

func (t *TagSpecificationSet) SetItems(items []TagSpecificationSetItem) {
    t.Items = items
}
