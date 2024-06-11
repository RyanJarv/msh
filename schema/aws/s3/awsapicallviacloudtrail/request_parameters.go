package awsapicallviacloudtrail

type RequestParameters struct {
    BucketName string `json:"bucketName"`
    LegalHold string `json:"legal-hold,omitempty"`
    Key string `json:"key"`
    Retention string `json:"retention,omitempty"`
    Host string `json:"Host"`
}

func (r *RequestParameters) SetBucketName(bucketName string) {
    r.BucketName = bucketName
}

func (r *RequestParameters) SetLegalHold(legalHold string) {
    r.LegalHold = legalHold
}

func (r *RequestParameters) SetKey(key string) {
    r.Key = key
}

func (r *RequestParameters) SetRetention(retention string) {
    r.Retention = retention
}

func (r *RequestParameters) SetHost(host string) {
    r.Host = host
}
