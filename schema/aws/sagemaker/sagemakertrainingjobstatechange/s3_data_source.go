package sagemakertrainingjobstatechange

type S3DataSource struct {
    S3DataDistributionType string `json:"S3DataDistributionType,omitempty"`
    S3Uri string `json:"S3Uri,omitempty"`
    S3DataType string `json:"S3DataType,omitempty"`
}

func (s *S3DataSource) SetS3DataDistributionType(s3DataDistributionType string) {
    s.S3DataDistributionType = s3DataDistributionType
}

func (s *S3DataSource) SetS3Uri(s3Uri string) {
    s.S3Uri = s3Uri
}

func (s *S3DataSource) SetS3DataType(s3DataType string) {
    s.S3DataType = s3DataType
}
