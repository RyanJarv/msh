package guarddutyfinding

type EbsVolumeDetails struct {
    ScannedVolumeDetails []EbsVolumeDetailsItem `json:"scannedVolumeDetails,omitempty"`
    SkippedVolumeDetails interface{} `json:"skippedVolumeDetails,omitempty"`
}

func (e *EbsVolumeDetails) SetScannedVolumeDetails(scannedVolumeDetails []EbsVolumeDetailsItem) {
    e.ScannedVolumeDetails = scannedVolumeDetails
}

func (e *EbsVolumeDetails) SetSkippedVolumeDetails(skippedVolumeDetails interface{}) {
    e.SkippedVolumeDetails = skippedVolumeDetails
}
