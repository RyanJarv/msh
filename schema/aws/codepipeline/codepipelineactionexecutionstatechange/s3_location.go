package codepipelineactionexecutionstatechange

type S3Location struct {
    Bucket string `json:"bucket"`
    Key string `json:"key"`
}

func (s *S3Location) SetBucket(bucket string) {
    s.Bucket = bucket
}

func (s *S3Location) SetKey(key string) {
    s.Key = key
}
