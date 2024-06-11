package guarddutyfinding

type ScannedItemCount struct {
    Files float64 `json:"files,omitempty"`
    TotalGb float64 `json:"totalGb,omitempty"`
    Volumes float64 `json:"volumes,omitempty"`
}

func (s *ScannedItemCount) SetFiles(files float64) {
    s.Files = files
}

func (s *ScannedItemCount) SetTotalGb(totalGb float64) {
    s.TotalGb = totalGb
}

func (s *ScannedItemCount) SetVolumes(volumes float64) {
    s.Volumes = volumes
}
