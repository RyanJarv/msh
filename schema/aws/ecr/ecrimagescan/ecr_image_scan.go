package ecrimagescan

type ECRImageScan struct {
    ImageDigest string `json:"image-digest,omitempty"`
    ImageTags []string `json:"image-tags,omitempty"`
    RepositoryName string `json:"repository-name"`
    ScanStatus string `json:"scan-status"`
    FindingSeverityCounts FindingSeverityCounts `json:"finding-severity-counts,omitempty"`
}

func (e *ECRImageScan) SetImageDigest(imageDigest string) {
    e.ImageDigest = imageDigest
}

func (e *ECRImageScan) SetImageTags(imageTags []string) {
    e.ImageTags = imageTags
}

func (e *ECRImageScan) SetRepositoryName(repositoryName string) {
    e.RepositoryName = repositoryName
}

func (e *ECRImageScan) SetScanStatus(scanStatus string) {
    e.ScanStatus = scanStatus
}

func (e *ECRImageScan) SetFindingSeverityCounts(findingSeverityCounts FindingSeverityCounts) {
    e.FindingSeverityCounts = findingSeverityCounts
}
