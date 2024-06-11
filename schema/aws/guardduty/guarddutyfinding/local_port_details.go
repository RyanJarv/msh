package guarddutyfinding

type LocalPortDetails struct {
    Port float64 `json:"port,omitempty"`
    PortName string `json:"portName,omitempty"`
}

func (l *LocalPortDetails) SetPort(port float64) {
    l.Port = port
}

func (l *LocalPortDetails) SetPortName(portName string) {
    l.PortName = portName
}
