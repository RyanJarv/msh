package awsapicallviacloudtrail

type NetworkInterfaceSet_1 struct {
    Items []NetworkInterfaceSet_1Item `json:"items,omitempty"`
}

func (n *NetworkInterfaceSet_1) SetItems(items []NetworkInterfaceSet_1Item) {
    n.Items = items
}
