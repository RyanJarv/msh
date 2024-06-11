package guarddutyfinding

type AccessKeyDetails struct {
    AccessKeyId string `json:"accessKeyId,omitempty"`
    PrincipalId string `json:"principalId,omitempty"`
    UserName string `json:"userName,omitempty"`
    UserType string `json:"userType,omitempty"`
}

func (a *AccessKeyDetails) SetAccessKeyId(accessKeyId string) {
    a.AccessKeyId = accessKeyId
}

func (a *AccessKeyDetails) SetPrincipalId(principalId string) {
    a.PrincipalId = principalId
}

func (a *AccessKeyDetails) SetUserName(userName string) {
    a.UserName = userName
}

func (a *AccessKeyDetails) SetUserType(userType string) {
    a.UserType = userType
}
