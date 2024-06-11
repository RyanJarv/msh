package awsapicallviacloudtrail

type TransformOutput struct {
    S3OutputPath string `json:"s3OutputPath,omitempty"`
}

func (t *TransformOutput) SetS3OutputPath(s3OutputPath string) {
    t.S3OutputPath = s3OutputPath
}
