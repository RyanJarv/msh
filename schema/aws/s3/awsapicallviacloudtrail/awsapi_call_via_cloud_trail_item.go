package awsapicallviacloudtrail

type AWSAPICallViaCloudTrailItem struct {
    ARN string `json:"ARN"`
    AccountId string `json:"accountId"`
    AWSAPICallViaCloudTrailItemType string `json:"type"`
}

func (a *AWSAPICallViaCloudTrailItem) SetARN(aRN string) {
    a.ARN = aRN
}

func (a *AWSAPICallViaCloudTrailItem) SetAccountId(accountId string) {
    a.AccountId = accountId
}

func (a *AWSAPICallViaCloudTrailItem) SetAWSAPICallViaCloudTrailItemType(aWSAPICallViaCloudTrailItemType string) {
    a.AWSAPICallViaCloudTrailItemType = aWSAPICallViaCloudTrailItemType
}
