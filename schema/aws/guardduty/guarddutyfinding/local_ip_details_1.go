package guarddutyfinding

type LocalIpDetails_1 struct {
    IpAddressV4 string `json:"ipAddressV4"`
}

func (l *LocalIpDetails_1) SetIpAddressV4(ipAddressV4 string) {
    l.IpAddressV4 = ipAddressV4
}
