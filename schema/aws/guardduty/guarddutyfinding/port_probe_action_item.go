package guarddutyfinding

type PortProbeActionItem struct {
    LocalIpDetails LocalIpDetails_1 `json:"localIpDetails"`
    LocalPortDetails LocalPortDetails_1 `json:"localPortDetails"`
    RemoteIpDetails RemoteIpDetails_4 `json:"remoteIpDetails"`
}

func (p *PortProbeActionItem) SetLocalIpDetails(localIpDetails LocalIpDetails_1) {
    p.LocalIpDetails = localIpDetails
}

func (p *PortProbeActionItem) SetLocalPortDetails(localPortDetails LocalPortDetails_1) {
    p.LocalPortDetails = localPortDetails
}

func (p *PortProbeActionItem) SetRemoteIpDetails(remoteIpDetails RemoteIpDetails_4) {
    p.RemoteIpDetails = remoteIpDetails
}
