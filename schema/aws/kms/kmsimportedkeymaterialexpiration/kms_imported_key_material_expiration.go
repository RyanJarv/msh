package kmsimportedkeymaterialexpiration

type KMSImportedKeyMaterialExpiration struct {
    KeyId string `json:"key-id"`
}

func (k *KMSImportedKeyMaterialExpiration) SetKeyId(keyId string) {
    k.KeyId = keyId
}
