package awsapicallviacloudtrail

type TransformResources struct {
    InstanceCount float64 `json:"instanceCount,omitempty"`
    InstanceType string `json:"instanceType,omitempty"`
}

func (t *TransformResources) SetInstanceCount(instanceCount float64) {
    t.InstanceCount = instanceCount
}

func (t *TransformResources) SetInstanceType(instanceType string) {
    t.InstanceType = instanceType
}
