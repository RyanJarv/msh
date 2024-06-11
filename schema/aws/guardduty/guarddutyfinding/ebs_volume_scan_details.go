package guarddutyfinding

type EbsVolumeScanDetails struct {
    ScanDetections ScanDetections `json:"scanDetections,omitempty"`
    ScanCompletedAt float64 `json:"scanCompletedAt,omitempty"`
    ScanId string `json:"scanId,omitempty"`
    ScanStartedAt float64 `json:"scanStartedAt,omitempty"`
    Sources []string `json:"sources,omitempty"`
    TriggerFindingId string `json:"triggerFindingId,omitempty"`
}

func (e *EbsVolumeScanDetails) SetScanDetections(scanDetections ScanDetections) {
    e.ScanDetections = scanDetections
}

func (e *EbsVolumeScanDetails) SetScanCompletedAt(scanCompletedAt float64) {
    e.ScanCompletedAt = scanCompletedAt
}

func (e *EbsVolumeScanDetails) SetScanId(scanId string) {
    e.ScanId = scanId
}

func (e *EbsVolumeScanDetails) SetScanStartedAt(scanStartedAt float64) {
    e.ScanStartedAt = scanStartedAt
}

func (e *EbsVolumeScanDetails) SetSources(sources []string) {
    e.Sources = sources
}

func (e *EbsVolumeScanDetails) SetTriggerFindingId(triggerFindingId string) {
    e.TriggerFindingId = triggerFindingId
}
