package guarddutyfinding

type LocalIpDetails struct {
    IpAddressV4 string `json:"ipAddressV4,omitempty"`
}

func (l *LocalIpDetails) SetIpAddressV4(ipAddressV4 string) {
    l.IpAddressV4 = ipAddressV4
}
