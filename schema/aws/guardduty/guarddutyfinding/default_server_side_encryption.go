package guarddutyfinding

type DefaultServerSideEncryption struct {
    EncryptionType string `json:"encryptionType,omitempty"`
    KmsMasterKeyArn string `json:"kmsMasterKeyArn,omitempty"`
}

func (d *DefaultServerSideEncryption) SetEncryptionType(encryptionType string) {
    d.EncryptionType = encryptionType
}

func (d *DefaultServerSideEncryption) SetKmsMasterKeyArn(kmsMasterKeyArn string) {
    d.KmsMasterKeyArn = kmsMasterKeyArn
}
