package guarddutyfinding

type KubernetesUserDetails struct {
    Groups []string `json:"groups,omitempty"`
    Uid string `json:"uid,omitempty"`
    Username string `json:"username,omitempty"`
}

func (k *KubernetesUserDetails) SetGroups(groups []string) {
    k.Groups = groups
}

func (k *KubernetesUserDetails) SetUid(uid string) {
    k.Uid = uid
}

func (k *KubernetesUserDetails) SetUsername(username string) {
    k.Username = username
}
