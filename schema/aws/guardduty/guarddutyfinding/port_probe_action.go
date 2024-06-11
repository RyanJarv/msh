package guarddutyfinding

type PortProbeAction struct {
    Blocked bool `json:"blocked,omitempty"`
    PortProbeDetails []PortProbeActionItem `json:"portProbeDetails,omitempty"`
}

func (p *PortProbeAction) SetBlocked(blocked bool) {
    p.Blocked = blocked
}

func (p *PortProbeAction) SetPortProbeDetails(portProbeDetails []PortProbeActionItem) {
    p.PortProbeDetails = portProbeDetails
}
