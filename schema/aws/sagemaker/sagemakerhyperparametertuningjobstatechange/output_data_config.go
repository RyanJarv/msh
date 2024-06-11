package sagemakerhyperparametertuningjobstatechange

type OutputDataConfig struct {
    KmsKeyId string `json:"KmsKeyId"`
    S3OutputPath string `json:"S3OutputPath"`
}

func (o *OutputDataConfig) SetKmsKeyId(kmsKeyId string) {
    o.KmsKeyId = kmsKeyId
}

func (o *OutputDataConfig) SetS3OutputPath(s3OutputPath string) {
    o.S3OutputPath = s3OutputPath
}
