package awsapicallviacloudtrail

type PrivateIpAddressesSet_1Item struct {
    PrivateIpAddress string `json:"privateIpAddress"`
    Primary bool `json:"primary"`
}

func (p *PrivateIpAddressesSet_1Item) SetPrivateIpAddress(privateIpAddress string) {
    p.PrivateIpAddress = privateIpAddress
}

func (p *PrivateIpAddressesSet_1Item) SetPrimary(primary bool) {
    p.Primary = primary
}
