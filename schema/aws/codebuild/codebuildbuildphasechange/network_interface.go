package codebuildbuildphasechange

type Network_interface struct {
    EniId string `json:"eni-id,omitempty"`
    SubnetId string `json:"subnet-id,omitempty"`
}

func (n *Network_interface) SetEniId(eniId string) {
    n.EniId = eniId
}

func (n *Network_interface) SetSubnetId(subnetId string) {
    n.SubnetId = subnetId
}
