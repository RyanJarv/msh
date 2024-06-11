package guarddutyfinding

type EvidenceItem struct {
    ThreatListName string `json:"threatListName"`
    ThreatNames []string `json:"threatNames"`
}

func (e *EvidenceItem) SetThreatListName(threatListName string) {
    e.ThreatListName = threatListName
}

func (e *EvidenceItem) SetThreatNames(threatNames []string) {
    e.ThreatNames = threatNames
}
