package guarddutyfinding

type ThreatDetectedByNameItem struct {
    FilePaths []ThreatDetectedByNameItemItem `json:"filePaths"`
    ItemCount float64 `json:"itemCount"`
    Name string `json:"name"`
    Severity string `json:"severity"`
}

func (t *ThreatDetectedByNameItem) SetFilePaths(filePaths []ThreatDetectedByNameItemItem) {
    t.FilePaths = filePaths
}

func (t *ThreatDetectedByNameItem) SetItemCount(itemCount float64) {
    t.ItemCount = itemCount
}

func (t *ThreatDetectedByNameItem) SetName(name string) {
    t.Name = name
}

func (t *ThreatDetectedByNameItem) SetSeverity(severity string) {
    t.Severity = severity
}
