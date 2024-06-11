package awsapicallviacloudtrail

type ExportSchemaRequestParameters struct {
    RegistryName string `json:"registryName"`
    SchemaName string `json:"schemaName"`
    SchemaVersion string `json:"schemaVersion,omitempty"`
    ExportSchemaRequestParametersType string `json:"type,omitempty"`
}

func (e *ExportSchemaRequestParameters) SetRegistryName(registryName string) {
    e.RegistryName = registryName
}

func (e *ExportSchemaRequestParameters) SetSchemaName(schemaName string) {
    e.SchemaName = schemaName
}

func (e *ExportSchemaRequestParameters) SetSchemaVersion(schemaVersion string) {
    e.SchemaVersion = schemaVersion
}

func (e *ExportSchemaRequestParameters) SetExportSchemaRequestParametersType(exportSchemaRequestParametersType string) {
    e.ExportSchemaRequestParametersType = exportSchemaRequestParametersType
}
