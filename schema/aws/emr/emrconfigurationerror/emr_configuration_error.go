package emrconfigurationerror

type EMRConfigurationError struct {
    Severity string `json:"severity,omitempty"`
    Description string `json:"description,omitempty"`
    ClusterId string `json:"clusterId,omitempty"`
}

func (e *EMRConfigurationError) SetSeverity(severity string) {
    e.Severity = severity
}

func (e *EMRConfigurationError) SetDescription(description string) {
    e.Description = description
}

func (e *EMRConfigurationError) SetClusterId(clusterId string) {
    e.ClusterId = clusterId
}
