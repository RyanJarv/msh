package awsapicallviacloudtrail

type NetworkInterfaceSet struct {
    Items []NetworkInterfaceSetItem `json:"items,omitempty"`
}

func (n *NetworkInterfaceSet) SetItems(items []NetworkInterfaceSetItem) {
    n.Items = items
}
