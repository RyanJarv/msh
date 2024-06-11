package awsapicallviacloudtrail

type NetworkInterface struct {
    GroupSet GroupSet_2 `json:"groupSet,omitempty"`
    TagSet interface{} `json:"tagSet,omitempty"`
    Ipv6AddressesSet interface{} `json:"ipv6AddressesSet,omitempty"`
    PrivateIpAddressesSet PrivateIpAddressesSet_1 `json:"privateIpAddressesSet,omitempty"`
    NetworkInterfaceId string `json:"networkInterfaceId,omitempty"`
    SubnetId string `json:"subnetId,omitempty"`
    RequesterId string `json:"requesterId,omitempty"`
    Description string `json:"description,omitempty"`
    OwnerId string `json:"ownerId,omitempty"`
    SourceDestCheck bool `json:"sourceDestCheck,omitempty"`
    AvailabilityZone string `json:"availabilityZone,omitempty"`
    RequesterManaged bool `json:"requesterManaged,omitempty"`
    PrivateIpAddress string `json:"privateIpAddress,omitempty"`
    InterfaceType string `json:"interfaceType,omitempty"`
    MacAddress string `json:"macAddress,omitempty"`
    VpcId string `json:"vpcId,omitempty"`
    Status string `json:"status,omitempty"`
}

func (n *NetworkInterface) SetGroupSet(groupSet GroupSet_2) {
    n.GroupSet = groupSet
}

func (n *NetworkInterface) SetTagSet(tagSet interface{}) {
    n.TagSet = tagSet
}

func (n *NetworkInterface) SetIpv6AddressesSet(ipv6AddressesSet interface{}) {
    n.Ipv6AddressesSet = ipv6AddressesSet
}

func (n *NetworkInterface) SetPrivateIpAddressesSet(privateIpAddressesSet PrivateIpAddressesSet_1) {
    n.PrivateIpAddressesSet = privateIpAddressesSet
}

func (n *NetworkInterface) SetNetworkInterfaceId(networkInterfaceId string) {
    n.NetworkInterfaceId = networkInterfaceId
}

func (n *NetworkInterface) SetSubnetId(subnetId string) {
    n.SubnetId = subnetId
}

func (n *NetworkInterface) SetRequesterId(requesterId string) {
    n.RequesterId = requesterId
}

func (n *NetworkInterface) SetDescription(description string) {
    n.Description = description
}

func (n *NetworkInterface) SetOwnerId(ownerId string) {
    n.OwnerId = ownerId
}

func (n *NetworkInterface) SetSourceDestCheck(sourceDestCheck bool) {
    n.SourceDestCheck = sourceDestCheck
}

func (n *NetworkInterface) SetAvailabilityZone(availabilityZone string) {
    n.AvailabilityZone = availabilityZone
}

func (n *NetworkInterface) SetRequesterManaged(requesterManaged bool) {
    n.RequesterManaged = requesterManaged
}

func (n *NetworkInterface) SetPrivateIpAddress(privateIpAddress string) {
    n.PrivateIpAddress = privateIpAddress
}

func (n *NetworkInterface) SetInterfaceType(interfaceType string) {
    n.InterfaceType = interfaceType
}

func (n *NetworkInterface) SetMacAddress(macAddress string) {
    n.MacAddress = macAddress
}

func (n *NetworkInterface) SetVpcId(vpcId string) {
    n.VpcId = vpcId
}

func (n *NetworkInterface) SetStatus(status string) {
    n.Status = status
}
