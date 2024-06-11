package awsapicallviacloudtrail

type UserIdentity struct {
    SessionContext SessionContext `json:"sessionContext"`
    AccessKeyId string `json:"accessKeyId"`
    AccountId string `json:"accountId"`
    Arn string `json:"arn"`
    PrincipalId string `json:"principalId"`
    UserIdentityType string `json:"type"`
}

func (u *UserIdentity) SetSessionContext(sessionContext SessionContext) {
    u.SessionContext = sessionContext
}

func (u *UserIdentity) SetAccessKeyId(accessKeyId string) {
    u.AccessKeyId = accessKeyId
}

func (u *UserIdentity) SetAccountId(accountId string) {
    u.AccountId = accountId
}

func (u *UserIdentity) SetArn(arn string) {
    u.Arn = arn
}

func (u *UserIdentity) SetPrincipalId(principalId string) {
    u.PrincipalId = principalId
}

func (u *UserIdentity) SetUserIdentityType(userIdentityType string) {
    u.UserIdentityType = userIdentityType
}
