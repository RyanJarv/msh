package awsapicallviacloudtrail

type DestinationS3Location struct {
    Bucket string `json:"bucket,omitempty"`
    Prefix string `json:"prefix,omitempty"`
}

func (d *DestinationS3Location) SetBucket(bucket string) {
    d.Bucket = bucket
}

func (d *DestinationS3Location) SetPrefix(prefix string) {
    d.Prefix = prefix
}
