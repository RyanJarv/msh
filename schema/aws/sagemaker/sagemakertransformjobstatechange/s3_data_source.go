package sagemakertransformjobstatechange

type S3DataSource struct {
    S3Uri string `json:"S3Uri"`
    S3DataType string `json:"S3DataType"`
}

func (s *S3DataSource) SetS3Uri(s3Uri string) {
    s.S3Uri = s3Uri
}

func (s *S3DataSource) SetS3DataType(s3DataType string) {
    s.S3DataType = s3DataType
}
