package awsapicallviacloudtrail

type DescribeSchemaRequestParameters struct {
    RegistryName string `json:"registryName"`
    SchemaName string `json:"schemaName"`
    SchemaVersion string `json:"schemaVersion,omitempty"`
}

func (d *DescribeSchemaRequestParameters) SetRegistryName(registryName string) {
    d.RegistryName = registryName
}

func (d *DescribeSchemaRequestParameters) SetSchemaName(schemaName string) {
    d.SchemaName = schemaName
}

func (d *DescribeSchemaRequestParameters) SetSchemaVersion(schemaVersion string) {
    d.SchemaVersion = schemaVersion
}
