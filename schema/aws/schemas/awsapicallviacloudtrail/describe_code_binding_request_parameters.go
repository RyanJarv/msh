package awsapicallviacloudtrail

type DescribeCodeBindingRequestParameters struct {
    RegistryName string `json:"registryName"`
    SchemaName string `json:"schemaName"`
    Language string `json:"language"`
    SchemaVersion string `json:"schemaVersion,omitempty"`
}

func (d *DescribeCodeBindingRequestParameters) SetRegistryName(registryName string) {
    d.RegistryName = registryName
}

func (d *DescribeCodeBindingRequestParameters) SetSchemaName(schemaName string) {
    d.SchemaName = schemaName
}

func (d *DescribeCodeBindingRequestParameters) SetLanguage(language string) {
    d.Language = language
}

func (d *DescribeCodeBindingRequestParameters) SetSchemaVersion(schemaVersion string) {
    d.SchemaVersion = schemaVersion
}
