package maciealert

type Object struct {
    PublicBucketFinancialSummaryTxt float64 `json:"public_bucket/financial_summary.txt,omitempty"`
    PublicBucketSecretKeyTxt float64 `json:"public_bucket/secret_key.txt,omitempty"`
}

func (o *Object) SetPublicBucketFinancialSummaryTxt(publicBucketFinancialSummaryTxt float64) {
    o.PublicBucketFinancialSummaryTxt = publicBucketFinancialSummaryTxt
}

func (o *Object) SetPublicBucketSecretKeyTxt(publicBucketSecretKeyTxt float64) {
    o.PublicBucketSecretKeyTxt = publicBucketSecretKeyTxt
}
