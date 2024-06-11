package awsapicallviacloudtrail

type SourceS3Location struct {
    Bucket string `json:"bucket,omitempty"`
    Key string `json:"key,omitempty"`
}

func (s *SourceS3Location) SetBucket(bucket string) {
    s.Bucket = bucket
}

func (s *SourceS3Location) SetKey(key string) {
    s.Key = key
}
