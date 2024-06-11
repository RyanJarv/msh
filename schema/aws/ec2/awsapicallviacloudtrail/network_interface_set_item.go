package awsapicallviacloudtrail

type NetworkInterfaceSetItem struct {
    SubnetId string `json:"subnetId"`
    DeviceIndex float64 `json:"deviceIndex"`
}

func (n *NetworkInterfaceSetItem) SetSubnetId(subnetId string) {
    n.SubnetId = subnetId
}

func (n *NetworkInterfaceSetItem) SetDeviceIndex(deviceIndex float64) {
    n.DeviceIndex = deviceIndex
}
