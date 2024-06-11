package guarddutyfinding

type InstanceDetailsItem struct {
    Ipv6Addresses []interface{} `json:"ipv6Addresses"`
    NetworkInterfaceId string `json:"networkInterfaceId"`
    PrivateDnsName string `json:"privateDnsName"`
    PrivateIpAddress string `json:"privateIpAddress"`
    PrivateIpAddresses []InstanceDetailsItemItem `json:"privateIpAddresses"`
    PublicDnsName string `json:"publicDnsName"`
    PublicIp string `json:"publicIp"`
    SecurityGroups []InstanceDetailsItemItem_1 `json:"securityGroups"`
    SubnetId string `json:"subnetId"`
    VpcId string `json:"vpcId"`
}

func (i *InstanceDetailsItem) SetIpv6Addresses(ipv6Addresses []interface{}) {
    i.Ipv6Addresses = ipv6Addresses
}

func (i *InstanceDetailsItem) SetNetworkInterfaceId(networkInterfaceId string) {
    i.NetworkInterfaceId = networkInterfaceId
}

func (i *InstanceDetailsItem) SetPrivateDnsName(privateDnsName string) {
    i.PrivateDnsName = privateDnsName
}

func (i *InstanceDetailsItem) SetPrivateIpAddress(privateIpAddress string) {
    i.PrivateIpAddress = privateIpAddress
}

func (i *InstanceDetailsItem) SetPrivateIpAddresses(privateIpAddresses []InstanceDetailsItemItem) {
    i.PrivateIpAddresses = privateIpAddresses
}

func (i *InstanceDetailsItem) SetPublicDnsName(publicDnsName string) {
    i.PublicDnsName = publicDnsName
}

func (i *InstanceDetailsItem) SetPublicIp(publicIp string) {
    i.PublicIp = publicIp
}

func (i *InstanceDetailsItem) SetSecurityGroups(securityGroups []InstanceDetailsItemItem_1) {
    i.SecurityGroups = securityGroups
}

func (i *InstanceDetailsItem) SetSubnetId(subnetId string) {
    i.SubnetId = subnetId
}

func (i *InstanceDetailsItem) SetVpcId(vpcId string) {
    i.VpcId = vpcId
}
