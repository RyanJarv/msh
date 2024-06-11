package awsapicallviacloudtrail

type PrivateIpAddressesSet_2 struct {
    Item []PrivateIpAddressesSet_1Item `json:"item"`
}

func (p *PrivateIpAddressesSet_2) SetItem(item []PrivateIpAddressesSet_1Item) {
    p.Item = item
}
