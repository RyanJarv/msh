package awsapicallviacloudtrail

type DeleteSchemaRequestParameters struct {
    RegistryName string `json:"registryName"`
    SchemaName string `json:"schemaName"`
}

func (d *DeleteSchemaRequestParameters) SetRegistryName(registryName string) {
    d.RegistryName = registryName
}

func (d *DeleteSchemaRequestParameters) SetSchemaName(schemaName string) {
    d.SchemaName = schemaName
}
