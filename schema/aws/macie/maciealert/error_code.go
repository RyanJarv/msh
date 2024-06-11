package maciealert

type Error_Code struct {
    NoSuchBucketPolicy float64 `json:"NoSuchBucketPolicy,omitempty"`
}

func (e *Error_Code) SetNoSuchBucketPolicy(noSuchBucketPolicy float64) {
    e.NoSuchBucketPolicy = noSuchBucketPolicy
}
