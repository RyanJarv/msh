package guarddutyfinding

type RemotePortDetails struct {
    Port float64 `json:"port,omitempty"`
    PortName string `json:"portName,omitempty"`
}

func (r *RemotePortDetails) SetPort(port float64) {
    r.Port = port
}

func (r *RemotePortDetails) SetPortName(portName string) {
    r.PortName = portName
}
