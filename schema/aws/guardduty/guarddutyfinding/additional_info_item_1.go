package guarddutyfinding

type AdditionalInfoItem_1 struct {
    AccessKeyId string `json:"accessKeyId"`
    IpAddressV4 string `json:"ipAddressV4"`
    PrincipalId string `json:"principalId"`
}

func (a *AdditionalInfoItem_1) SetAccessKeyId(accessKeyId string) {
    a.AccessKeyId = accessKeyId
}

func (a *AdditionalInfoItem_1) SetIpAddressV4(ipAddressV4 string) {
    a.IpAddressV4 = ipAddressV4
}

func (a *AdditionalInfoItem_1) SetPrincipalId(principalId string) {
    a.PrincipalId = principalId
}
