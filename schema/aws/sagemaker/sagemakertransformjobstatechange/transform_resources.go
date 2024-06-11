package sagemakertransformjobstatechange

type TransformResources struct {
    InstanceCount float64 `json:"InstanceCount"`
    InstanceType string `json:"InstanceType"`
}

func (t *TransformResources) SetInstanceCount(instanceCount float64) {
    t.InstanceCount = instanceCount
}

func (t *TransformResources) SetInstanceType(instanceType string) {
    t.InstanceType = instanceType
}
