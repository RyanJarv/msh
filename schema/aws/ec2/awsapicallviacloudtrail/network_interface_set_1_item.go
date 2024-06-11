package awsapicallviacloudtrail

type NetworkInterfaceSet_1Item struct {
    GroupSet GroupSet_3 `json:"groupSet"`
    TagSet interface{} `json:"tagSet"`
    Attachment Attachment `json:"attachment"`
    Ipv6AddressesSet interface{} `json:"ipv6AddressesSet"`
    PrivateIpAddressesSet PrivateIpAddressesSet_2 `json:"privateIpAddressesSet"`
    NetworkInterfaceId string `json:"networkInterfaceId"`
    SubnetId string `json:"subnetId"`
    OwnerId string `json:"ownerId"`
    SourceDestCheck bool `json:"sourceDestCheck"`
    PrivateIpAddress string `json:"privateIpAddress"`
    InterfaceType string `json:"interfaceType"`
    MacAddress string `json:"macAddress"`
    VpcId string `json:"vpcId"`
    Status string `json:"status"`
}

func (n *NetworkInterfaceSet_1Item) SetGroupSet(groupSet GroupSet_3) {
    n.GroupSet = groupSet
}

func (n *NetworkInterfaceSet_1Item) SetTagSet(tagSet interface{}) {
    n.TagSet = tagSet
}

func (n *NetworkInterfaceSet_1Item) SetAttachment(attachment Attachment) {
    n.Attachment = attachment
}

func (n *NetworkInterfaceSet_1Item) SetIpv6AddressesSet(ipv6AddressesSet interface{}) {
    n.Ipv6AddressesSet = ipv6AddressesSet
}

func (n *NetworkInterfaceSet_1Item) SetPrivateIpAddressesSet(privateIpAddressesSet PrivateIpAddressesSet_2) {
    n.PrivateIpAddressesSet = privateIpAddressesSet
}

func (n *NetworkInterfaceSet_1Item) SetNetworkInterfaceId(networkInterfaceId string) {
    n.NetworkInterfaceId = networkInterfaceId
}

func (n *NetworkInterfaceSet_1Item) SetSubnetId(subnetId string) {
    n.SubnetId = subnetId
}

func (n *NetworkInterfaceSet_1Item) SetOwnerId(ownerId string) {
    n.OwnerId = ownerId
}

func (n *NetworkInterfaceSet_1Item) SetSourceDestCheck(sourceDestCheck bool) {
    n.SourceDestCheck = sourceDestCheck
}

func (n *NetworkInterfaceSet_1Item) SetPrivateIpAddress(privateIpAddress string) {
    n.PrivateIpAddress = privateIpAddress
}

func (n *NetworkInterfaceSet_1Item) SetInterfaceType(interfaceType string) {
    n.InterfaceType = interfaceType
}

func (n *NetworkInterfaceSet_1Item) SetMacAddress(macAddress string) {
    n.MacAddress = macAddress
}

func (n *NetworkInterfaceSet_1Item) SetVpcId(vpcId string) {
    n.VpcId = vpcId
}

func (n *NetworkInterfaceSet_1Item) SetStatus(status string) {
    n.Status = status
}
