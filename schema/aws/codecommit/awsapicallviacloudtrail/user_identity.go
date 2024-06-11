package awsapicallviacloudtrail

type UserIdentity struct {
    SessionContext SessionContext `json:"sessionContext,omitempty"`
    AccessKeyId string `json:"accessKeyId,omitempty"`
    AccountId string `json:"accountId"`
    Arn string `json:"arn"`
    PrincipalId string `json:"principalId"`
    UserIdentityType string `json:"type"`
    UserName string `json:"userName,omitempty"`
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

func (u *UserIdentity) SetUserName(userName string) {
    u.UserName = userName
}
