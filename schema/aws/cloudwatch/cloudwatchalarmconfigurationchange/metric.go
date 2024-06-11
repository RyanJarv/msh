package cloudwatchalarmconfigurationchange

type Metric struct {
    Dimensions interface{} `json:"dimensions"`
    Namespace string `json:"namespace"`
    Name string `json:"name"`
}

func (m *Metric) SetDimensions(dimensions interface{}) {
    m.Dimensions = dimensions
}

func (m *Metric) SetNamespace(namespace string) {
    m.Namespace = namespace
}

func (m *Metric) SetName(name string) {
    m.Name = name
}
