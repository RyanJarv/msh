package guarddutyfinding

type ScanDetections struct {
    HighestSeverityThreatDetails HighestSeverityThreatDetails `json:"highestSeverityThreatDetails,omitempty"`
    ScannedItemCount ScannedItemCount `json:"scannedItemCount,omitempty"`
    ThreatDetectedByName ThreatDetectedByName `json:"threatDetectedByName,omitempty"`
    ThreatsDetectedItemCount ThreatsDetectedItemCount `json:"threatsDetectedItemCount,omitempty"`
}

func (s *ScanDetections) SetHighestSeverityThreatDetails(highestSeverityThreatDetails HighestSeverityThreatDetails) {
    s.HighestSeverityThreatDetails = highestSeverityThreatDetails
}

func (s *ScanDetections) SetScannedItemCount(scannedItemCount ScannedItemCount) {
    s.ScannedItemCount = scannedItemCount
}

func (s *ScanDetections) SetThreatDetectedByName(threatDetectedByName ThreatDetectedByName) {
    s.ThreatDetectedByName = threatDetectedByName
}

func (s *ScanDetections) SetThreatsDetectedItemCount(threatsDetectedItemCount ThreatsDetectedItemCount) {
    s.ThreatsDetectedItemCount = threatsDetectedItemCount
}
