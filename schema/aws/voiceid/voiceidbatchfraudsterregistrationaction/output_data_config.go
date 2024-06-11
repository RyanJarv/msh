package voiceidbatchfraudsterregistrationaction

type OutputDataConfig struct {
    KmsKeyId string `json:"kmsKeyId,omitempty"`
    S3Uri string `json:"s3Uri,omitempty"`
}

func (o *OutputDataConfig) SetKmsKeyId(kmsKeyId string) {
    o.KmsKeyId = kmsKeyId
}

func (o *OutputDataConfig) SetS3Uri(s3Uri string) {
    o.S3Uri = s3Uri
}
