package awsapicallviacloudtrail

type AWSAPICallViaCloudTrailItem struct {
    ARN string `json:"ARN,omitempty"`
    AccountId string `json:"accountId,omitempty"`
    AWSAPICallViaCloudTrailItemType string `json:"type,omitempty"`
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
