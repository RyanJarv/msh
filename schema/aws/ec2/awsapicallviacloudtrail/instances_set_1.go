package awsapicallviacloudtrail

type InstancesSet_1 struct {
    Items []InstancesSet_1Item `json:"items,omitempty"`
}

func (i *InstancesSet_1) SetItems(items []InstancesSet_1Item) {
    i.Items = items
}
