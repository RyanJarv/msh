package awsapicallviacloudtrail

type S3DataSource struct {
    S3Uri string `json:"s3Uri,omitempty"`
    S3DataType string `json:"s3DataType,omitempty"`
}

func (s *S3DataSource) SetS3Uri(s3Uri string) {
    s.S3Uri = s3Uri
}

func (s *S3DataSource) SetS3DataType(s3DataType string) {
    s.S3DataType = s3DataType
}
