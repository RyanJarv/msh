package guarddutyfinding

type KubernetesWorkloadDetailsItem struct {
    SecurityContext SecurityContext `json:"securityContext"`
    Image string `json:"image"`
    ImagePrefix string `json:"imagePrefix"`
    Name string `json:"name"`
}

func (k *KubernetesWorkloadDetailsItem) SetSecurityContext(securityContext SecurityContext) {
    k.SecurityContext = securityContext
}

func (k *KubernetesWorkloadDetailsItem) SetImage(image string) {
    k.Image = image
}

func (k *KubernetesWorkloadDetailsItem) SetImagePrefix(imagePrefix string) {
    k.ImagePrefix = imagePrefix
}

func (k *KubernetesWorkloadDetailsItem) SetName(name string) {
    k.Name = name
}
