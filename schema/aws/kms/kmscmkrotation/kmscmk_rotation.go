package kmscmkrotation

type KMSCMKRotation struct {
    KeyId string `json:"key-id"`
}

func (k *KMSCMKRotation) SetKeyId(keyId string) {
    k.KeyId = keyId
}
