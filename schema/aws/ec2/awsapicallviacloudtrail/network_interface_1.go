package awsapicallviacloudtrail

type NetworkInterface_1 struct {
    SecurityGroupId SecurityGroupId `json:"SecurityGroupId,omitempty"`
    DeviceIndex float64 `json:"DeviceIndex,omitempty"`
    Tag float64 `json:"tag,omitempty"`
    SubnetId string `json:"SubnetId,omitempty"`
}

func (n *NetworkInterface_1) SetSecurityGroupId(securityGroupId SecurityGroupId) {
    n.SecurityGroupId = securityGroupId
}

func (n *NetworkInterface_1) SetDeviceIndex(deviceIndex float64) {
    n.DeviceIndex = deviceIndex
}

func (n *NetworkInterface_1) SetTag(tag float64) {
    n.Tag = tag
}

func (n *NetworkInterface_1) SetSubnetId(subnetId string) {
    n.SubnetId = subnetId
}
