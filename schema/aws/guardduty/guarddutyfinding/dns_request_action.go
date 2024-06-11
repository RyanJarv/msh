package guarddutyfinding

type DnsRequestAction struct {
    Blocked bool `json:"blocked,omitempty"`
    Domain string `json:"domain,omitempty"`
    Protocol string `json:"protocol,omitempty"`
}

func (d *DnsRequestAction) SetBlocked(blocked bool) {
    d.Blocked = blocked
}

func (d *DnsRequestAction) SetDomain(domain string) {
    d.Domain = domain
}

func (d *DnsRequestAction) SetProtocol(protocol string) {
    d.Protocol = protocol
}
