package maciealert

type DistinctEventName struct {
    Narrative string `json:"narrative,omitempty"`
    Name string `json:"name,omitempty"`
    Description string `json:"description,omitempty"`
    Risk float64 `json:"risk,omitempty"`
}

func (d *DistinctEventName) SetNarrative(narrative string) {
    d.Narrative = narrative
}

func (d *DistinctEventName) SetName(name string) {
    d.Name = name
}

func (d *DistinctEventName) SetDescription(description string) {
    d.Description = description
}

func (d *DistinctEventName) SetRisk(risk float64) {
    d.Risk = risk
}
