package awsapicallviacloudtrail

type OutputDataConfig struct {
    RemoveJobNameFromS3OutputPath bool `json:"removeJobNameFromS3OutputPath,omitempty"`
    S3OutputPath string `json:"s3OutputPath,omitempty"`
}

func (o *OutputDataConfig) SetRemoveJobNameFromS3OutputPath(removeJobNameFromS3OutputPath bool) {
    o.RemoveJobNameFromS3OutputPath = removeJobNameFromS3OutputPath
}

func (o *OutputDataConfig) SetS3OutputPath(s3OutputPath string) {
    o.S3OutputPath = s3OutputPath
}
