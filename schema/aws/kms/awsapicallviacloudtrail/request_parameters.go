package awsapicallviacloudtrail

type RequestParameters struct {
    EncryptionContext EncryptionContext `json:"encryptionContext"`
    KeyId string `json:"keyId"`
    KeySpec string `json:"keySpec"`
}

func (r *RequestParameters) SetEncryptionContext(encryptionContext EncryptionContext) {
    r.EncryptionContext = encryptionContext
}

func (r *RequestParameters) SetKeyId(keyId string) {
    r.KeyId = keyId
}

func (r *RequestParameters) SetKeySpec(keySpec string) {
    r.KeySpec = keySpec
}
