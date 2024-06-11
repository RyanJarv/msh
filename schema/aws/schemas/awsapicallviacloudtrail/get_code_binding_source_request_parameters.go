package awsapicallviacloudtrail

type GetCodeBindingSourceRequestParameters struct {
    RegistryName string `json:"registryName"`
    SchemaName string `json:"schemaName"`
    Language string `json:"language"`
}

func (g *GetCodeBindingSourceRequestParameters) SetRegistryName(registryName string) {
    g.RegistryName = registryName
}

func (g *GetCodeBindingSourceRequestParameters) SetSchemaName(schemaName string) {
    g.SchemaName = schemaName
}

func (g *GetCodeBindingSourceRequestParameters) SetLanguage(language string) {
    g.Language = language
}
