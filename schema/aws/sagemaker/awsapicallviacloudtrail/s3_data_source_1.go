package awsapicallviacloudtrail

type S3DataSource_1 struct {
    S3Uri string `json:"s3Uri"`
    S3DataType string `json:"s3DataType"`
    S3DataDistributionType string `json:"s3DataDistributionType"`
}

func (s *S3DataSource_1) SetS3Uri(s3Uri string) {
    s.S3Uri = s3Uri
}

func (s *S3DataSource_1) SetS3DataType(s3DataType string) {
    s.S3DataType = s3DataType
}

func (s *S3DataSource_1) SetS3DataDistributionType(s3DataDistributionType string) {
    s.S3DataDistributionType = s3DataDistributionType
}
