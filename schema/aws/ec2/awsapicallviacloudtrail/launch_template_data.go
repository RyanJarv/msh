package awsapicallviacloudtrail

type LaunchTemplateData struct {
    InstanceMarketOptions InstanceMarketOptions_1 `json:"InstanceMarketOptions,omitempty"`
    NetworkInterface NetworkInterface_1 `json:"NetworkInterface,omitempty"`
    UserData string `json:"UserData,omitempty"`
    ImageId string `json:"ImageId,omitempty"`
    InstanceType string `json:"InstanceType,omitempty"`
}

func (l *LaunchTemplateData) SetInstanceMarketOptions(instanceMarketOptions InstanceMarketOptions_1) {
    l.InstanceMarketOptions = instanceMarketOptions
}

func (l *LaunchTemplateData) SetNetworkInterface(networkInterface NetworkInterface_1) {
    l.NetworkInterface = networkInterface
}

func (l *LaunchTemplateData) SetUserData(userData string) {
    l.UserData = userData
}

func (l *LaunchTemplateData) SetImageId(imageId string) {
    l.ImageId = imageId
}

func (l *LaunchTemplateData) SetInstanceType(instanceType string) {
    l.InstanceType = instanceType
}
