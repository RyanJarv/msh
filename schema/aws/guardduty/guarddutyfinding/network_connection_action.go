package guarddutyfinding

type NetworkConnectionAction struct {
    LocalIpDetails LocalIpDetails `json:"localIpDetails,omitempty"`
    LocalPortDetails LocalPortDetails `json:"localPortDetails,omitempty"`
    RemoteIpDetails RemoteIpDetails_3 `json:"remoteIpDetails,omitempty"`
    RemotePortDetails RemotePortDetails `json:"remotePortDetails,omitempty"`
    Blocked bool `json:"blocked,omitempty"`
    ConnectionDirection string `json:"connectionDirection,omitempty"`
    Protocol string `json:"protocol,omitempty"`
}

func (n *NetworkConnectionAction) SetLocalIpDetails(localIpDetails LocalIpDetails) {
    n.LocalIpDetails = localIpDetails
}

func (n *NetworkConnectionAction) SetLocalPortDetails(localPortDetails LocalPortDetails) {
    n.LocalPortDetails = localPortDetails
}

func (n *NetworkConnectionAction) SetRemoteIpDetails(remoteIpDetails RemoteIpDetails_3) {
    n.RemoteIpDetails = remoteIpDetails
}

func (n *NetworkConnectionAction) SetRemotePortDetails(remotePortDetails RemotePortDetails) {
    n.RemotePortDetails = remotePortDetails
}

func (n *NetworkConnectionAction) SetBlocked(blocked bool) {
    n.Blocked = blocked
}

func (n *NetworkConnectionAction) SetConnectionDirection(connectionDirection string) {
    n.ConnectionDirection = connectionDirection
}

func (n *NetworkConnectionAction) SetProtocol(protocol string) {
    n.Protocol = protocol
}
