package simpleworkflowexecutionstatechange

type WorkflowType struct {
    Name string `json:"name"`
    Version string `json:"version"`
}

func (w *WorkflowType) SetName(name string) {
    w.Name = name
}

func (w *WorkflowType) SetVersion(version string) {
    w.Version = version
}
