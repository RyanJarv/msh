package awsapicallviacloudtrail

type AdditionalEventData struct {
    ObjectRetentionInfo ObjectRetentionInfo `json:"objectRetentionInfo,omitempty"`
    XAmzId2 string `json:"x-amz-id-2"`
    AuthenticationMethod string `json:"AuthenticationMethod"`
    CipherSuite string `json:"CipherSuite"`
    SignatureVersion string `json:"SignatureVersion"`
    BytesTransferredIn float64 `json:"bytesTransferredIn"`
    BytesTransferredOut float64 `json:"bytesTransferredOut"`
}

func (a *AdditionalEventData) SetObjectRetentionInfo(objectRetentionInfo ObjectRetentionInfo) {
    a.ObjectRetentionInfo = objectRetentionInfo
}

func (a *AdditionalEventData) SetXAmzId2(xAmzId2 string) {
    a.XAmzId2 = xAmzId2
}

func (a *AdditionalEventData) SetAuthenticationMethod(authenticationMethod string) {
    a.AuthenticationMethod = authenticationMethod
}

func (a *AdditionalEventData) SetCipherSuite(cipherSuite string) {
    a.CipherSuite = cipherSuite
}

func (a *AdditionalEventData) SetSignatureVersion(signatureVersion string) {
    a.SignatureVersion = signatureVersion
}

func (a *AdditionalEventData) SetBytesTransferredIn(bytesTransferredIn float64) {
    a.BytesTransferredIn = bytesTransferredIn
}

func (a *AdditionalEventData) SetBytesTransferredOut(bytesTransferredOut float64) {
    a.BytesTransferredOut = bytesTransferredOut
}
