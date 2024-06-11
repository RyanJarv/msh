package awsapicallviacloudtrail

type AWSAPICallViaCloudTrailItem struct {
    AccountId string `json:"accountId"`
    ARN string `json:"ARN"`
}

func (a *AWSAPICallViaCloudTrailItem) SetAccountId(accountId string) {
    a.AccountId = accountId
}

func (a *AWSAPICallViaCloudTrailItem) SetARN(aRN string) {
    a.ARN = aRN
}
