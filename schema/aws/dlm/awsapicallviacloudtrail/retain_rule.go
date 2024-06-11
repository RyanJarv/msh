package awsapicallviacloudtrail

type RetainRule struct {
    Count float64 `json:"Count"`
}

func (r *RetainRule) SetCount(count float64) {
    r.Count = count
}
