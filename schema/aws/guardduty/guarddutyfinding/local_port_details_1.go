package guarddutyfinding

type LocalPortDetails_1 struct {
    Port float64 `json:"port"`
    PortName string `json:"portName"`
}

func (l *LocalPortDetails_1) SetPort(port float64) {
    l.Port = port
}

func (l *LocalPortDetails_1) SetPortName(portName string) {
    l.PortName = portName
}
