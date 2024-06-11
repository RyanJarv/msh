package awsapicallviacloudtrail

type DeleteSchemaVersionRequestParameters struct {
    RegistryName string `json:"registryName"`
    SchemaName string `json:"schemaName"`
    SchemaVersion string `json:"schemaVersion"`
}

func (d *DeleteSchemaVersionRequestParameters) SetRegistryName(registryName string) {
    d.RegistryName = registryName
}

func (d *DeleteSchemaVersionRequestParameters) SetSchemaName(schemaName string) {
    d.SchemaName = schemaName
}

func (d *DeleteSchemaVersionRequestParameters) SetSchemaVersion(schemaVersion string) {
    d.SchemaVersion = schemaVersion
}
