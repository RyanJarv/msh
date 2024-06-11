package awsxrayinsightupdate

type RootCauseServiceId struct {
    AccountId string `json:"AccountId"`
    Name string `json:"Name"`
    Names []interface{} `json:"Names"`
    RootCauseServiceIdType string `json:"Type"`
}

func (r *RootCauseServiceId) SetAccountId(accountId string) {
    r.AccountId = accountId
}

func (r *RootCauseServiceId) SetName(name string) {
    r.Name = name
}

func (r *RootCauseServiceId) SetNames(names []interface{}) {
    r.Names = names
}

func (r *RootCauseServiceId) SetRootCauseServiceIdType(rootCauseServiceIdType string) {
    r.RootCauseServiceIdType = rootCauseServiceIdType
}
