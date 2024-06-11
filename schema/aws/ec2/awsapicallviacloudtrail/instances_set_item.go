package awsapicallviacloudtrail

type InstancesSetItem struct {
    CapacityReservationSpecification CapacityReservationSpecification `json:"capacityReservationSpecification,omitempty"`
    StateReason StateReason `json:"stateReason,omitempty"`
    GroupSet GroupSet_2 `json:"groupSet,omitempty"`
    TagSet TagSet `json:"tagSet,omitempty"`
    InstanceState InstanceState `json:"instanceState,omitempty"`
    ProductCodes interface{} `json:"productCodes,omitempty"`
    NetworkInterfaceSet NetworkInterfaceSet_1 `json:"networkInterfaceSet,omitempty"`
    BlockDeviceMapping interface{} `json:"blockDeviceMapping,omitempty"`
    CpuOptions CpuOptions `json:"cpuOptions,omitempty"`
    Monitoring Monitoring_1 `json:"monitoring,omitempty"`
    PreviousState InstanceState `json:"previousState,omitempty"`
    EnclaveOptions EnclaveOptions `json:"enclaveOptions,omitempty"`
    Placement Placement `json:"placement,omitempty"`
    CurrentState InstanceState `json:"currentState,omitempty"`
    SubnetId string `json:"subnetId,omitempty"`
    VirtualizationType string `json:"virtualizationType,omitempty"`
    AmiLaunchIndex float64 `json:"amiLaunchIndex,omitempty"`
    SourceDestCheck bool `json:"sourceDestCheck,omitempty"`
    InstanceId string `json:"instanceId"`
    VpcId string `json:"vpcId,omitempty"`
    Hypervisor string `json:"hypervisor,omitempty"`
    RootDeviceName string `json:"rootDeviceName,omitempty"`
    Architecture string `json:"architecture,omitempty"`
    EbsOptimized bool `json:"ebsOptimized,omitempty"`
    ImageId string `json:"imageId,omitempty"`
    ClientToken string `json:"clientToken,omitempty"`
    InstanceType string `json:"instanceType,omitempty"`
    PrivateIpAddress string `json:"privateIpAddress,omitempty"`
    InstanceLifecycle string `json:"instanceLifecycle,omitempty"`
    RootDeviceType string `json:"rootDeviceType,omitempty"`
    LaunchTime int64 `json:"launchTime,omitempty"`
    SpotInstanceRequestId string `json:"spotInstanceRequestId,omitempty"`
}

func (i *InstancesSetItem) SetCapacityReservationSpecification(capacityReservationSpecification CapacityReservationSpecification) {
    i.CapacityReservationSpecification = capacityReservationSpecification
}

func (i *InstancesSetItem) SetStateReason(stateReason StateReason) {
    i.StateReason = stateReason
}

func (i *InstancesSetItem) SetGroupSet(groupSet GroupSet_2) {
    i.GroupSet = groupSet
}

func (i *InstancesSetItem) SetTagSet(tagSet TagSet) {
    i.TagSet = tagSet
}

func (i *InstancesSetItem) SetInstanceState(instanceState InstanceState) {
    i.InstanceState = instanceState
}

func (i *InstancesSetItem) SetProductCodes(productCodes interface{}) {
    i.ProductCodes = productCodes
}

func (i *InstancesSetItem) SetNetworkInterfaceSet(networkInterfaceSet NetworkInterfaceSet_1) {
    i.NetworkInterfaceSet = networkInterfaceSet
}

func (i *InstancesSetItem) SetBlockDeviceMapping(blockDeviceMapping interface{}) {
    i.BlockDeviceMapping = blockDeviceMapping
}

func (i *InstancesSetItem) SetCpuOptions(cpuOptions CpuOptions) {
    i.CpuOptions = cpuOptions
}

func (i *InstancesSetItem) SetMonitoring(monitoring Monitoring_1) {
    i.Monitoring = monitoring
}

func (i *InstancesSetItem) SetPreviousState(previousState InstanceState) {
    i.PreviousState = previousState
}

func (i *InstancesSetItem) SetEnclaveOptions(enclaveOptions EnclaveOptions) {
    i.EnclaveOptions = enclaveOptions
}

func (i *InstancesSetItem) SetPlacement(placement Placement) {
    i.Placement = placement
}

func (i *InstancesSetItem) SetCurrentState(currentState InstanceState) {
    i.CurrentState = currentState
}

func (i *InstancesSetItem) SetSubnetId(subnetId string) {
    i.SubnetId = subnetId
}

func (i *InstancesSetItem) SetVirtualizationType(virtualizationType string) {
    i.VirtualizationType = virtualizationType
}

func (i *InstancesSetItem) SetAmiLaunchIndex(amiLaunchIndex float64) {
    i.AmiLaunchIndex = amiLaunchIndex
}

func (i *InstancesSetItem) SetSourceDestCheck(sourceDestCheck bool) {
    i.SourceDestCheck = sourceDestCheck
}

func (i *InstancesSetItem) SetInstanceId(instanceId string) {
    i.InstanceId = instanceId
}

func (i *InstancesSetItem) SetVpcId(vpcId string) {
    i.VpcId = vpcId
}

func (i *InstancesSetItem) SetHypervisor(hypervisor string) {
    i.Hypervisor = hypervisor
}

func (i *InstancesSetItem) SetRootDeviceName(rootDeviceName string) {
    i.RootDeviceName = rootDeviceName
}

func (i *InstancesSetItem) SetArchitecture(architecture string) {
    i.Architecture = architecture
}

func (i *InstancesSetItem) SetEbsOptimized(ebsOptimized bool) {
    i.EbsOptimized = ebsOptimized
}

func (i *InstancesSetItem) SetImageId(imageId string) {
    i.ImageId = imageId
}

func (i *InstancesSetItem) SetClientToken(clientToken string) {
    i.ClientToken = clientToken
}

func (i *InstancesSetItem) SetInstanceType(instanceType string) {
    i.InstanceType = instanceType
}

func (i *InstancesSetItem) SetPrivateIpAddress(privateIpAddress string) {
    i.PrivateIpAddress = privateIpAddress
}

func (i *InstancesSetItem) SetInstanceLifecycle(instanceLifecycle string) {
    i.InstanceLifecycle = instanceLifecycle
}

func (i *InstancesSetItem) SetRootDeviceType(rootDeviceType string) {
    i.RootDeviceType = rootDeviceType
}

func (i *InstancesSetItem) SetLaunchTime(launchTime int64) {
    i.LaunchTime = launchTime
}

func (i *InstancesSetItem) SetSpotInstanceRequestId(spotInstanceRequestId string) {
    i.SpotInstanceRequestId = spotInstanceRequestId
}
