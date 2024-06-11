package guarddutyfinding

import (
    "time"
)


type InstanceDetails struct {
    IamInstanceProfile IamInstanceProfile `json:"iamInstanceProfile,omitempty"`
    AvailabilityZone string `json:"availabilityZone,omitempty"`
    ImageDescription string `json:"imageDescription,omitempty"`
    ImageId string `json:"imageId,omitempty"`
    InstanceId string `json:"instanceId,omitempty"`
    InstanceState string `json:"instanceState,omitempty"`
    InstanceType string `json:"instanceType,omitempty"`
    LaunchTime time.Time `json:"launchTime,omitempty"`
    NetworkInterfaces []InstanceDetailsItem `json:"networkInterfaces,omitempty"`
    OutpostArn string `json:"outpostArn,omitempty"`
    Platform interface{} `json:"platform,omitempty"`
    ProductCodes []InstanceDetailsItem_1 `json:"productCodes,omitempty"`
    Tags []EcsClusterDetailsItem `json:"tags,omitempty"`
}

func (i *InstanceDetails) SetIamInstanceProfile(iamInstanceProfile IamInstanceProfile) {
    i.IamInstanceProfile = iamInstanceProfile
}

func (i *InstanceDetails) SetAvailabilityZone(availabilityZone string) {
    i.AvailabilityZone = availabilityZone
}

func (i *InstanceDetails) SetImageDescription(imageDescription string) {
    i.ImageDescription = imageDescription
}

func (i *InstanceDetails) SetImageId(imageId string) {
    i.ImageId = imageId
}

func (i *InstanceDetails) SetInstanceId(instanceId string) {
    i.InstanceId = instanceId
}

func (i *InstanceDetails) SetInstanceState(instanceState string) {
    i.InstanceState = instanceState
}

func (i *InstanceDetails) SetInstanceType(instanceType string) {
    i.InstanceType = instanceType
}

func (i *InstanceDetails) SetLaunchTime(launchTime time.Time) {
    i.LaunchTime = launchTime
}

func (i *InstanceDetails) SetNetworkInterfaces(networkInterfaces []InstanceDetailsItem) {
    i.NetworkInterfaces = networkInterfaces
}

func (i *InstanceDetails) SetOutpostArn(outpostArn string) {
    i.OutpostArn = outpostArn
}

func (i *InstanceDetails) SetPlatform(platform interface{}) {
    i.Platform = platform
}

func (i *InstanceDetails) SetProductCodes(productCodes []InstanceDetailsItem_1) {
    i.ProductCodes = productCodes
}

func (i *InstanceDetails) SetTags(tags []EcsClusterDetailsItem) {
    i.Tags = tags
}
