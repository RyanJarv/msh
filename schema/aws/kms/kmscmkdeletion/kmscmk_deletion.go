package kmscmkdeletion

type KMSCMKDeletion struct {
    KeyId string `json:"key-id"`
}

func (k *KMSCMKDeletion) SetKeyId(keyId string) {
    k.KeyId = keyId
}
