package sagemakertrainingjobstatechange

type ResourceConfig struct {
    InstanceCount float64 `json:"InstanceCount,omitempty"`
    VolumeSizeInGB float64 `json:"VolumeSizeInGB,omitempty"`
    InstanceType string `json:"InstanceType,omitempty"`
}

func (r *ResourceConfig) SetInstanceCount(instanceCount float64) {
    r.InstanceCount = instanceCount
}

func (r *ResourceConfig) SetVolumeSizeInGB(volumeSizeInGB float64) {
    r.VolumeSizeInGB = volumeSizeInGB
}

func (r *ResourceConfig) SetInstanceType(instanceType string) {
    r.InstanceType = instanceType
}
