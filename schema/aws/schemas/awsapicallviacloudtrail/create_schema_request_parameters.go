package awsapicallviacloudtrail

type CreateSchemaRequestParameters struct {
    RegistryName string `json:"registryName"`
    SchemaName string `json:"schemaName"`
    CreateSchemaRequestParametersType string `json:"Type"`
    Description string `json:"Description,omitempty"`
    Content string `json:"Content"`
    Tags Tags `json:"tags,omitempty"`
}

func (c *CreateSchemaRequestParameters) SetRegistryName(registryName string) {
    c.RegistryName = registryName
}

func (c *CreateSchemaRequestParameters) SetSchemaName(schemaName string) {
    c.SchemaName = schemaName
}

func (c *CreateSchemaRequestParameters) SetCreateSchemaRequestParametersType(createSchemaRequestParametersType string) {
    c.CreateSchemaRequestParametersType = createSchemaRequestParametersType
}

func (c *CreateSchemaRequestParameters) SetDescription(description string) {
    c.Description = description
}

func (c *CreateSchemaRequestParameters) SetContent(content string) {
    c.Content = content
}

func (c *CreateSchemaRequestParameters) SetTags(tags Tags) {
    c.Tags = tags
}
