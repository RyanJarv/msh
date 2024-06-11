package awsapicallviacloudtrail

type ResourceConfig struct {
    VolumeSizeInGB float64 `json:"volumeSizeInGB,omitempty"`
    InstanceCount float64 `json:"instanceCount,omitempty"`
    InstanceType string `json:"instanceType,omitempty"`
}

func (r *ResourceConfig) SetVolumeSizeInGB(volumeSizeInGB float64) {
    r.VolumeSizeInGB = volumeSizeInGB
}

func (r *ResourceConfig) SetInstanceCount(instanceCount float64) {
    r.InstanceCount = instanceCount
}

func (r *ResourceConfig) SetInstanceType(instanceType string) {
    r.InstanceType = instanceType
}
