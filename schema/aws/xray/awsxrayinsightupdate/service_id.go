package awsxrayinsightupdate

type ServiceId struct {
    AccountId string `json:"AccountId"`
    Name string `json:"Name"`
    Names []interface{} `json:"Names"`
    ServiceIdType string `json:"Type"`
}

func (s *ServiceId) SetAccountId(accountId string) {
    s.AccountId = accountId
}

func (s *ServiceId) SetName(name string) {
    s.Name = name
}

func (s *ServiceId) SetNames(names []interface{}) {
    s.Names = names
}

func (s *ServiceId) SetServiceIdType(serviceIdType string) {
    s.ServiceIdType = serviceIdType
}
