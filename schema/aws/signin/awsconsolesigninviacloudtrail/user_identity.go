package awsconsolesigninviacloudtrail

type UserIdentity struct {
    AccountId string `json:"accountId"`
    PrincipalId string `json:"principalId"`
    UserIdentityType string `json:"type"`
    Arn string `json:"arn"`
}

func (u *UserIdentity) SetAccountId(accountId string) {
    u.AccountId = accountId
}

func (u *UserIdentity) SetPrincipalId(principalId string) {
    u.PrincipalId = principalId
}

func (u *UserIdentity) SetUserIdentityType(userIdentityType string) {
    u.UserIdentityType = userIdentityType
}

func (u *UserIdentity) SetArn(arn string) {
    u.Arn = arn
}
