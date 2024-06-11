package awsapicallviacloudtrail

type ListSchemaVersionRequestParameters struct {
    RegistryName string `json:"registryName"`
    Limit string `json:"limit,omitempty"`
    NextToken string `json:"nextToken,omitempty"`
    SchemaName string `json:"schemaName"`
}

func (l *ListSchemaVersionRequestParameters) SetRegistryName(registryName string) {
    l.RegistryName = registryName
}

func (l *ListSchemaVersionRequestParameters) SetLimit(limit string) {
    l.Limit = limit
}

func (l *ListSchemaVersionRequestParameters) SetNextToken(nextToken string) {
    l.NextToken = nextToken
}

func (l *ListSchemaVersionRequestParameters) SetSchemaName(schemaName string) {
    l.SchemaName = schemaName
}
