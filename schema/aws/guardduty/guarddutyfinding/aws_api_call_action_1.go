package guarddutyfinding

type AwsApiCallAction_1 struct {
    AffectedResources AffectedResources_1 `json:"affectedResources,omitempty"`
    RemoteAccountDetails RemoteAccountDetails `json:"remoteAccountDetails,omitempty"`
    RemoteIpDetails RemoteIpDetails_1 `json:"remoteIpDetails,omitempty"`
    Api string `json:"api,omitempty"`
    CallerType string `json:"callerType,omitempty"`
    ErrorCode string `json:"errorCode,omitempty"`
    ServiceName string `json:"serviceName,omitempty"`
}

func (a *AwsApiCallAction_1) SetAffectedResources(affectedResources AffectedResources_1) {
    a.AffectedResources = affectedResources
}

func (a *AwsApiCallAction_1) SetRemoteAccountDetails(remoteAccountDetails RemoteAccountDetails) {
    a.RemoteAccountDetails = remoteAccountDetails
}

func (a *AwsApiCallAction_1) SetRemoteIpDetails(remoteIpDetails RemoteIpDetails_1) {
    a.RemoteIpDetails = remoteIpDetails
}

func (a *AwsApiCallAction_1) SetApi(api string) {
    a.Api = api
}

func (a *AwsApiCallAction_1) SetCallerType(callerType string) {
    a.CallerType = callerType
}

func (a *AwsApiCallAction_1) SetErrorCode(errorCode string) {
    a.ErrorCode = errorCode
}

func (a *AwsApiCallAction_1) SetServiceName(serviceName string) {
    a.ServiceName = serviceName
}
