package awsserviceeventviacloudtrail

type ServiceEventDetails struct {
    Response string `json:"response"`
}

func (s *ServiceEventDetails) SetResponse(response string) {
    s.Response = response
}
