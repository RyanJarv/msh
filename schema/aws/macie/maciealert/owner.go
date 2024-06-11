package maciealert

type Owner struct {
    BucketOwner float64 `json:"bucket_owner,omitempty"`
}

func (o *Owner) SetBucketOwner(bucketOwner float64) {
    o.BucketOwner = bucketOwner
}
