package guarddutyfinding

type InstanceDetailsItemItem struct {
    PrivateDnsName string `json:"privateDnsName"`
    PrivateIpAddress string `json:"privateIpAddress"`
}

func (i *InstanceDetailsItemItem) SetPrivateDnsName(privateDnsName string) {
    i.PrivateDnsName = privateDnsName
}

func (i *InstanceDetailsItemItem) SetPrivateIpAddress(privateIpAddress string) {
    i.PrivateIpAddress = privateIpAddress
}
