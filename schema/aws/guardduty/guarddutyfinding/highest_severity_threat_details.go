package guarddutyfinding

type HighestSeverityThreatDetails struct {
    Count float64 `json:"count,omitempty"`
    Severity string `json:"severity,omitempty"`
    ThreatName string `json:"threatName,omitempty"`
}

func (h *HighestSeverityThreatDetails) SetCount(count float64) {
    h.Count = count
}

func (h *HighestSeverityThreatDetails) SetSeverity(severity string) {
    h.Severity = severity
}

func (h *HighestSeverityThreatDetails) SetThreatName(threatName string) {
    h.ThreatName = threatName
}
