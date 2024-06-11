package awsserviceeventviacloudtrail

type ServiceEventDetails struct {
    CreateAccountStatus CreateAccountStatus `json:"createAccountStatus"`
}

func (s *ServiceEventDetails) SetCreateAccountStatus(createAccountStatus CreateAccountStatus) {
    s.CreateAccountStatus = createAccountStatus
}
