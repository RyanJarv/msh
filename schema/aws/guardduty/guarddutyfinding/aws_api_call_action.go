package guarddutyfinding

type AwsApiCallAction struct {
    AffectedResources AffectedResources `json:"affectedResources,omitempty"`
    RemoteIpDetails RemoteIpDetails `json:"remoteIpDetails,omitempty"`
    Api string `json:"api,omitempty"`
    CallerType string `json:"callerType,omitempty"`
    ErrorCode string `json:"errorCode,omitempty"`
    ServiceName string `json:"serviceName,omitempty"`
}

func (a *AwsApiCallAction) SetAffectedResources(affectedResources AffectedResources) {
    a.AffectedResources = affectedResources
}

func (a *AwsApiCallAction) SetRemoteIpDetails(remoteIpDetails RemoteIpDetails) {
    a.RemoteIpDetails = remoteIpDetails
}

func (a *AwsApiCallAction) SetApi(api string) {
    a.Api = api
}

func (a *AwsApiCallAction) SetCallerType(callerType string) {
    a.CallerType = callerType
}

func (a *AwsApiCallAction) SetErrorCode(errorCode string) {
    a.ErrorCode = errorCode
}

func (a *AwsApiCallAction) SetServiceName(serviceName string) {
    a.ServiceName = serviceName
}
