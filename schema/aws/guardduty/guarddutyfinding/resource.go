package guarddutyfinding

type Resource struct {
    AccessKeyDetails AccessKeyDetails `json:"accessKeyDetails,omitempty"`
    ContainerDetails ContainerDetails `json:"containerDetails,omitempty"`
    EbsVolumeDetails EbsVolumeDetails `json:"ebsVolumeDetails,omitempty"`
    EcsClusterDetails EcsClusterDetails `json:"ecsClusterDetails,omitempty"`
    EksClusterDetails EksClusterDetails `json:"eksClusterDetails,omitempty"`
    InstanceDetails InstanceDetails `json:"instanceDetails,omitempty"`
    KubernetesDetails KubernetesDetails `json:"kubernetesDetails,omitempty"`
    ResourceType string `json:"resourceType"`
    S3BucketDetails []ResourceItem `json:"s3BucketDetails,omitempty"`
}

func (r *Resource) SetAccessKeyDetails(accessKeyDetails AccessKeyDetails) {
    r.AccessKeyDetails = accessKeyDetails
}

func (r *Resource) SetContainerDetails(containerDetails ContainerDetails) {
    r.ContainerDetails = containerDetails
}

func (r *Resource) SetEbsVolumeDetails(ebsVolumeDetails EbsVolumeDetails) {
    r.EbsVolumeDetails = ebsVolumeDetails
}

func (r *Resource) SetEcsClusterDetails(ecsClusterDetails EcsClusterDetails) {
    r.EcsClusterDetails = ecsClusterDetails
}

func (r *Resource) SetEksClusterDetails(eksClusterDetails EksClusterDetails) {
    r.EksClusterDetails = eksClusterDetails
}

func (r *Resource) SetInstanceDetails(instanceDetails InstanceDetails) {
    r.InstanceDetails = instanceDetails
}

func (r *Resource) SetKubernetesDetails(kubernetesDetails KubernetesDetails) {
    r.KubernetesDetails = kubernetesDetails
}

func (r *Resource) SetResourceType(resourceType string) {
    r.ResourceType = resourceType
}

func (r *Resource) SetS3BucketDetails(s3BucketDetails []ResourceItem) {
    r.S3BucketDetails = s3BucketDetails
}
