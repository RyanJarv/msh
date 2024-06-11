package guarddutyfinding

type ThreatDetectedByName struct {
    ItemCount float64 `json:"itemCount,omitempty"`
    Shortened bool `json:"shortened,omitempty"`
    ThreatNames []ThreatDetectedByNameItem `json:"threatNames,omitempty"`
    UniqueThreatNameCount float64 `json:"uniqueThreatNameCount,omitempty"`
}

func (t *ThreatDetectedByName) SetItemCount(itemCount float64) {
    t.ItemCount = itemCount
}

func (t *ThreatDetectedByName) SetShortened(shortened bool) {
    t.Shortened = shortened
}

func (t *ThreatDetectedByName) SetThreatNames(threatNames []ThreatDetectedByNameItem) {
    t.ThreatNames = threatNames
}

func (t *ThreatDetectedByName) SetUniqueThreatNameCount(uniqueThreatNameCount float64) {
    t.UniqueThreatNameCount = uniqueThreatNameCount
}
