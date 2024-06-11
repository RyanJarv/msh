package awsapicallviacloudtrail

type PutCodeBindingRequestParameters struct {
    RegistryName string `json:"registryName"`
    SchemaName string `json:"schemaName"`
    Language string `json:"language"`
    SchemaVersion string `json:"schemaVersion,omitempty"`
}

func (p *PutCodeBindingRequestParameters) SetRegistryName(registryName string) {
    p.RegistryName = registryName
}

func (p *PutCodeBindingRequestParameters) SetSchemaName(schemaName string) {
    p.SchemaName = schemaName
}

func (p *PutCodeBindingRequestParameters) SetLanguage(language string) {
    p.Language = language
}

func (p *PutCodeBindingRequestParameters) SetSchemaVersion(schemaVersion string) {
    p.SchemaVersion = schemaVersion
}
