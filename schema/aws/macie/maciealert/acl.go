package maciealert

type ACL struct {
    SecretBucketName []ACLItem `json:"secret-bucket-name,omitempty"`
}

func (a *ACL) SetSecretBucketName(secretBucketName []ACLItem) {
    a.SecretBucketName = secretBucketName
}
