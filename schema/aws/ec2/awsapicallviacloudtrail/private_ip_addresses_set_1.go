package awsapicallviacloudtrail

type PrivateIpAddressesSet_1 struct {
    Item []PrivateIpAddressesSet_1Item `json:"item,omitempty"`
}

func (p *PrivateIpAddressesSet_1) SetItem(item []PrivateIpAddressesSet_1Item) {
    p.Item = item
}
