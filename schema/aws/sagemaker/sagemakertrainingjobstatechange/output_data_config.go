package sagemakertrainingjobstatechange

type OutputDataConfig struct {
    S3OutputPath string `json:"S3OutputPath,omitempty"`
}

func (o *OutputDataConfig) SetS3OutputPath(s3OutputPath string) {
    o.S3OutputPath = s3OutputPath
}
