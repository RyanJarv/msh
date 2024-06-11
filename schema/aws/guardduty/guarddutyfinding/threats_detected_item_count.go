package guarddutyfinding

type ThreatsDetectedItemCount struct {
    Files float64 `json:"files,omitempty"`
}

func (t *ThreatsDetectedItemCount) SetFiles(files float64) {
    t.Files = files
}
