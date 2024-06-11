package awsapicallviacloudtrail

type InstancesSet_1Item struct {
    InstanceId string `json:"instanceId,omitempty"`
    ImageId string `json:"imageId,omitempty"`
    MinCount float64 `json:"minCount,omitempty"`
    MaxCount float64 `json:"maxCount,omitempty"`
}

func (i *InstancesSet_1Item) SetInstanceId(instanceId string) {
    i.InstanceId = instanceId
}

func (i *InstancesSet_1Item) SetImageId(imageId string) {
    i.ImageId = imageId
}

func (i *InstancesSet_1Item) SetMinCount(minCount float64) {
    i.MinCount = minCount
}

func (i *InstancesSet_1Item) SetMaxCount(maxCount float64) {
    i.MaxCount = maxCount
}
