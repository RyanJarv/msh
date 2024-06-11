package guarddutyfinding

type KubernetesWorkloadDetails struct {
    Containers []KubernetesWorkloadDetailsItem `json:"containers,omitempty"`
    Name string `json:"name,omitempty"`
    Namespace string `json:"namespace,omitempty"`
    KubernetesWorkloadDetailsType string `json:"type,omitempty"`
    Uid string `json:"uid,omitempty"`
}

func (k *KubernetesWorkloadDetails) SetContainers(containers []KubernetesWorkloadDetailsItem) {
    k.Containers = containers
}

func (k *KubernetesWorkloadDetails) SetName(name string) {
    k.Name = name
}

func (k *KubernetesWorkloadDetails) SetNamespace(namespace string) {
    k.Namespace = namespace
}

func (k *KubernetesWorkloadDetails) SetKubernetesWorkloadDetailsType(kubernetesWorkloadDetailsType string) {
    k.KubernetesWorkloadDetailsType = kubernetesWorkloadDetailsType
}

func (k *KubernetesWorkloadDetails) SetUid(uid string) {
    k.Uid = uid
}
