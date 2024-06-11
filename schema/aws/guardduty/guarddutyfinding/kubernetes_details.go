package guarddutyfinding

type KubernetesDetails struct {
    KubernetesUserDetails KubernetesUserDetails `json:"kubernetesUserDetails,omitempty"`
    KubernetesWorkloadDetails KubernetesWorkloadDetails `json:"kubernetesWorkloadDetails,omitempty"`
}

func (k *KubernetesDetails) SetKubernetesUserDetails(kubernetesUserDetails KubernetesUserDetails) {
    k.KubernetesUserDetails = kubernetesUserDetails
}

func (k *KubernetesDetails) SetKubernetesWorkloadDetails(kubernetesWorkloadDetails KubernetesWorkloadDetails) {
    k.KubernetesWorkloadDetails = kubernetesWorkloadDetails
}
