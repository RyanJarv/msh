package awsapicallviacloudtrail

type AWSAPICallViaCloudTrailItem struct {
    AccountId string `json:"accountId"`
    AWSAPICallViaCloudTrailItemType string `json:"type"`
    ARN string `json:"ARN"`
}

func (a *AWSAPICallViaCloudTrailItem) SetAccountId(accountId string) {
    a.AccountId = accountId
}

func (a *AWSAPICallViaCloudTrailItem) SetAWSAPICallViaCloudTrailItemType(aWSAPICallViaCloudTrailItemType string) {
    a.AWSAPICallViaCloudTrailItemType = aWSAPICallViaCloudTrailItemType
}

func (a *AWSAPICallViaCloudTrailItem) SetARN(aRN string) {
    a.ARN = aRN
}
