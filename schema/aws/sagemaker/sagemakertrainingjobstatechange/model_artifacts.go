package sagemakertrainingjobstatechange

type ModelArtifacts struct {
    S3ModelArtifacts string `json:"S3ModelArtifacts,omitempty"`
}

func (m *ModelArtifacts) SetS3ModelArtifacts(s3ModelArtifacts string) {
    m.S3ModelArtifacts = s3ModelArtifacts
}
