package guarddutyfinding

type RemoteAccountDetails struct {
    AccountId string `json:"accountId,omitempty"`
    Affiliated bool `json:"affiliated,omitempty"`
}

func (r *RemoteAccountDetails) SetAccountId(accountId string) {
    r.AccountId = accountId
}

func (r *RemoteAccountDetails) SetAffiliated(affiliated bool) {
    r.Affiliated = affiliated
}
