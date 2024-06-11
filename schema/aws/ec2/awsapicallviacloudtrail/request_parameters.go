package awsapicallviacloudtrail

type RequestParameters struct {
    CreateLaunchTemplateRequest CreateLaunchTemplateRequest `json:"CreateLaunchTemplateRequest,omitempty"`
    GroupSet GroupSet_1 `json:"groupSet,omitempty"`
    InstancesSet InstancesSet_1 `json:"instancesSet,omitempty"`
    CreateFleetRequest CreateFleetRequest `json:"CreateFleetRequest,omitempty"`
    DeleteLaunchTemplateRequest DeleteLaunchTemplateRequest `json:"DeleteLaunchTemplateRequest,omitempty"`
    PrivateIpAddressesSet interface{} `json:"privateIpAddressesSet,omitempty"`
    NetworkInterfaceSet NetworkInterfaceSet `json:"networkInterfaceSet,omitempty"`
    BlockDeviceMapping interface{} `json:"blockDeviceMapping,omitempty"`
    TagSpecificationSet TagSpecificationSet `json:"tagSpecificationSet,omitempty"`
    Monitoring Monitoring `json:"monitoring,omitempty"`
    InstanceMarketOptions InstanceMarketOptions `json:"instanceMarketOptions,omitempty"`
    LaunchTemplate LaunchTemplate `json:"launchTemplate,omitempty"`
    SubnetId string `json:"subnetId,omitempty"`
    UserData string `json:"userData,omitempty"`
    GroupId string `json:"groupId,omitempty"`
    Description string `json:"description,omitempty"`
    AvailabilityZone string `json:"availabilityZone,omitempty"`
    VpcId string `json:"vpcId,omitempty"`
    NetworkInterfaceId string `json:"networkInterfaceId,omitempty"`
    ClientToken string `json:"clientToken,omitempty"`
    InstanceType string `json:"instanceType,omitempty"`
    Ipv6AddressCount float64 `json:"ipv6AddressCount,omitempty"`
    GroupName string `json:"groupName,omitempty"`
    DisableApiTermination bool `json:"disableApiTermination,omitempty"`
    GroupDescription string `json:"groupDescription,omitempty"`
}

func (r *RequestParameters) SetCreateLaunchTemplateRequest(createLaunchTemplateRequest CreateLaunchTemplateRequest) {
    r.CreateLaunchTemplateRequest = createLaunchTemplateRequest
}

func (r *RequestParameters) SetGroupSet(groupSet GroupSet_1) {
    r.GroupSet = groupSet
}

func (r *RequestParameters) SetInstancesSet(instancesSet InstancesSet_1) {
    r.InstancesSet = instancesSet
}

func (r *RequestParameters) SetCreateFleetRequest(createFleetRequest CreateFleetRequest) {
    r.CreateFleetRequest = createFleetRequest
}

func (r *RequestParameters) SetDeleteLaunchTemplateRequest(deleteLaunchTemplateRequest DeleteLaunchTemplateRequest) {
    r.DeleteLaunchTemplateRequest = deleteLaunchTemplateRequest
}

func (r *RequestParameters) SetPrivateIpAddressesSet(privateIpAddressesSet interface{}) {
    r.PrivateIpAddressesSet = privateIpAddressesSet
}

func (r *RequestParameters) SetNetworkInterfaceSet(networkInterfaceSet NetworkInterfaceSet) {
    r.NetworkInterfaceSet = networkInterfaceSet
}

func (r *RequestParameters) SetBlockDeviceMapping(blockDeviceMapping interface{}) {
    r.BlockDeviceMapping = blockDeviceMapping
}

func (r *RequestParameters) SetTagSpecificationSet(tagSpecificationSet TagSpecificationSet) {
    r.TagSpecificationSet = tagSpecificationSet
}

func (r *RequestParameters) SetMonitoring(monitoring Monitoring) {
    r.Monitoring = monitoring
}

func (r *RequestParameters) SetInstanceMarketOptions(instanceMarketOptions InstanceMarketOptions) {
    r.InstanceMarketOptions = instanceMarketOptions
}

func (r *RequestParameters) SetLaunchTemplate(launchTemplate LaunchTemplate) {
    r.LaunchTemplate = launchTemplate
}

func (r *RequestParameters) SetSubnetId(subnetId string) {
    r.SubnetId = subnetId
}

func (r *RequestParameters) SetUserData(userData string) {
    r.UserData = userData
}

func (r *RequestParameters) SetGroupId(groupId string) {
    r.GroupId = groupId
}

func (r *RequestParameters) SetDescription(description string) {
    r.Description = description
}

func (r *RequestParameters) SetAvailabilityZone(availabilityZone string) {
    r.AvailabilityZone = availabilityZone
}

func (r *RequestParameters) SetVpcId(vpcId string) {
    r.VpcId = vpcId
}

func (r *RequestParameters) SetNetworkInterfaceId(networkInterfaceId string) {
    r.NetworkInterfaceId = networkInterfaceId
}

func (r *RequestParameters) SetClientToken(clientToken string) {
    r.ClientToken = clientToken
}

func (r *RequestParameters) SetInstanceType(instanceType string) {
    r.InstanceType = instanceType
}

func (r *RequestParameters) SetIpv6AddressCount(ipv6AddressCount float64) {
    r.Ipv6AddressCount = ipv6AddressCount
}

func (r *RequestParameters) SetGroupName(groupName string) {
    r.GroupName = groupName
}

func (r *RequestParameters) SetDisableApiTermination(disableApiTermination bool) {
    r.DisableApiTermination = disableApiTermination
}

func (r *RequestParameters) SetGroupDescription(groupDescription string) {
    r.GroupDescription = groupDescription
}
