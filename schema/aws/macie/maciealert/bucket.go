package maciealert

type Bucket struct {
    SecretBucketName float64 `json:"secret-bucket-name,omitempty"`
}

func (b *Bucket) SetSecretBucketName(secretBucketName float64) {
    b.SecretBucketName = secretBucketName
}
