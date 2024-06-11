package mediapackageharvestjobnotification

type S3_destination struct {
    BucketName string `json:"bucket_name"`
    ManifestKey string `json:"manifest_key"`
    RoleArn string `json:"role_arn"`
}

func (s *S3_destination) SetBucketName(bucketName string) {
    s.BucketName = bucketName
}

func (s *S3_destination) SetManifestKey(manifestKey string) {
    s.ManifestKey = manifestKey
}

func (s *S3_destination) SetRoleArn(roleArn string) {
    s.RoleArn = roleArn
}
