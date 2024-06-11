package awsapicallviacloudtrail

type InstancesSet struct {
    Items []InstancesSetItem `json:"items,omitempty"`
}

func (i *InstancesSet) SetItems(items []InstancesSetItem) {
    i.Items = items
}
