package awsserviceeventviacloudtrail

type CreateAccountStatus struct {
    AccountId string `json:"accountId"`
    AccountName string `json:"accountName"`
    CompletedTimestamp string `json:"completedTimestamp"`
    Id string `json:"id"`
    State string `json:"state"`
    RequestedTimestamp string `json:"requestedTimestamp"`
}

func (c *CreateAccountStatus) SetAccountId(accountId string) {
    c.AccountId = accountId
}

func (c *CreateAccountStatus) SetAccountName(accountName string) {
    c.AccountName = accountName
}

func (c *CreateAccountStatus) SetCompletedTimestamp(completedTimestamp string) {
    c.CompletedTimestamp = completedTimestamp
}

func (c *CreateAccountStatus) SetId(id string) {
    c.Id = id
}

func (c *CreateAccountStatus) SetState(state string) {
    c.State = state
}

func (c *CreateAccountStatus) SetRequestedTimestamp(requestedTimestamp string) {
    c.RequestedTimestamp = requestedTimestamp
}
