package awsserviceeventviacloudtrail

type UserIdentity struct {
    AccountId string `json:"accountId"`
    InvokedBy string `json:"invokedBy"`
}

func (u *UserIdentity) SetAccountId(accountId string) {
    u.AccountId = accountId
}

func (u *UserIdentity) SetInvokedBy(invokedBy string) {
    u.InvokedBy = invokedBy
}
