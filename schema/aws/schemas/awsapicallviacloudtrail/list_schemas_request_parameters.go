package awsapicallviacloudtrail

type ListSchemasRequestParameters struct {
    RegistryName string `json:"registryName"`
    Limit string `json:"limit,omitempty"`
    NextToken string `json:"nextToken,omitempty"`
    SchemaNamePrefix string `json:"schemaNamePrefix,omitempty"`
}

func (l *ListSchemasRequestParameters) SetRegistryName(registryName string) {
    l.RegistryName = registryName
}

func (l *ListSchemasRequestParameters) SetLimit(limit string) {
    l.Limit = limit
}

func (l *ListSchemasRequestParameters) SetNextToken(nextToken string) {
    l.NextToken = nextToken
}

func (l *ListSchemasRequestParameters) SetSchemaNamePrefix(schemaNamePrefix string) {
    l.SchemaNamePrefix = schemaNamePrefix
}
