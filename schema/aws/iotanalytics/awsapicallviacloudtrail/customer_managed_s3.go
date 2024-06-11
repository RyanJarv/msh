package awsapicallviacloudtrail

type CustomerManagedS3 struct {
    Bucket string `json:"bucket,omitempty"`
    KeyPrefix string `json:"keyPrefix,omitempty"`
    RoleArn string `json:"roleArn,omitempty"`
}

func (c *CustomerManagedS3) SetBucket(bucket string) {
    c.Bucket = bucket
}

func (c *CustomerManagedS3) SetKeyPrefix(keyPrefix string) {
    c.KeyPrefix = keyPrefix
}

func (c *CustomerManagedS3) SetRoleArn(roleArn string) {
    c.RoleArn = roleArn
}
