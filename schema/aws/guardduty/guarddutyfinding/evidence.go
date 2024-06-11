package guarddutyfinding

type Evidence struct {
    ThreatIntelligenceDetails []EvidenceItem `json:"threatIntelligenceDetails,omitempty"`
}

func (e *Evidence) SetThreatIntelligenceDetails(threatIntelligenceDetails []EvidenceItem) {
    e.ThreatIntelligenceDetails = threatIntelligenceDetails
}
