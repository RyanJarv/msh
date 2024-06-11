package awsapicallviacloudtrail

type Artifacts struct {
    EncryptionDisabled bool `json:"encryptionDisabled,omitempty"`
    Location string `json:"location,omitempty"`
}

func (a *Artifacts) SetEncryptionDisabled(encryptionDisabled bool) {
    a.EncryptionDisabled = encryptionDisabled
}

func (a *Artifacts) SetLocation(location string) {
    a.Location = location
}
