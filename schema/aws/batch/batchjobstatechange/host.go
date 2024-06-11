package batchjobstatechange

type Host struct {
    SourcePath string `json:"sourcePath,omitempty"`
}

func (h *Host) SetSourcePath(sourcePath string) {
    h.SourcePath = sourcePath
}
