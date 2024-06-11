package awsapicallviacloudtrail

type Artifacts_1 struct {
    EncryptionDisabled bool `json:"encryptionDisabled,omitempty"`
    Location string `json:"location,omitempty"`
    Name string `json:"name,omitempty"`
    NamespaceType string `json:"namespaceType,omitempty"`
    Packaging string `json:"packaging,omitempty"`
    Artifacts_1Type string `json:"type,omitempty"`
}

func (a *Artifacts_1) SetEncryptionDisabled(encryptionDisabled bool) {
    a.EncryptionDisabled = encryptionDisabled
}

func (a *Artifacts_1) SetLocation(location string) {
    a.Location = location
}

func (a *Artifacts_1) SetName(name string) {
    a.Name = name
}

func (a *Artifacts_1) SetNamespaceType(namespaceType string) {
    a.NamespaceType = namespaceType
}

func (a *Artifacts_1) SetPackaging(packaging string) {
    a.Packaging = packaging
}

func (a *Artifacts_1) SetArtifacts1Type(artifacts1Type string) {
    a.Artifacts_1Type = artifacts1Type
}
