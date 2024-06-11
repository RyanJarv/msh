package batchjobstatechange

type NetworkInterface struct {
    AttachmentId float64 `json:"attachmentId,omitempty"`
    Ipv6Address float64 `json:"ipv6Address,omitempty"`
    PrivateIpv4Address string `json:"privateIpv4Address,omitempty"`
}

func (n *NetworkInterface) SetAttachmentId(attachmentId float64) {
    n.AttachmentId = attachmentId
}

func (n *NetworkInterface) SetIpv6Address(ipv6Address float64) {
    n.Ipv6Address = ipv6Address
}

func (n *NetworkInterface) SetPrivateIpv4Address(privateIpv4Address string) {
    n.PrivateIpv4Address = privateIpv4Address
}
