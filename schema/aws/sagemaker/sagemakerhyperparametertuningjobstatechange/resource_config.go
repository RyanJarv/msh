package sagemakerhyperparametertuningjobstatechange

type ResourceConfig struct {
    InstanceCount float64 `json:"InstanceCount"`
    VolumeSizeInGB float64 `json:"VolumeSizeInGB"`
    VolumeKmsKeyId string `json:"VolumeKmsKeyId"`
    InstanceType string `json:"InstanceType"`
}

func (r *ResourceConfig) SetInstanceCount(instanceCount float64) {
    r.InstanceCount = instanceCount
}

func (r *ResourceConfig) SetVolumeSizeInGB(volumeSizeInGB float64) {
    r.VolumeSizeInGB = volumeSizeInGB
}

func (r *ResourceConfig) SetVolumeKmsKeyId(volumeKmsKeyId string) {
    r.VolumeKmsKeyId = volumeKmsKeyId
}

func (r *ResourceConfig) SetInstanceType(instanceType string) {
    r.InstanceType = instanceType
}
