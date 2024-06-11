package awsapicallviacloudtrail

type EncryptionContext struct {
    AwsS3Arn string `json:"aws:s3:arn"`
}

func (e *EncryptionContext) SetAwsS3Arn(awsS3Arn string) {
    e.AwsS3Arn = awsS3Arn
}
