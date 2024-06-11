package awsapicallviacloudtrail

type UpdateSchemaRequestParameters struct {
    UpdateSchemaRequestParametersType string `json:"Type,omitempty"`
    RegistryName string `json:"registryName"`
    SchemaName string `json:"schemaName"`
    Description string `json:"Description,omitempty"`
    ClientTokenId string `json:"ClientTokenId,omitempty"`
    Content string `json:"Content,omitempty"`
}

func (u *UpdateSchemaRequestParameters) SetUpdateSchemaRequestParametersType(updateSchemaRequestParametersType string) {
    u.UpdateSchemaRequestParametersType = updateSchemaRequestParametersType
}

func (u *UpdateSchemaRequestParameters) SetRegistryName(registryName string) {
    u.RegistryName = registryName
}

func (u *UpdateSchemaRequestParameters) SetSchemaName(schemaName string) {
    u.SchemaName = schemaName
}

func (u *UpdateSchemaRequestParameters) SetDescription(description string) {
    u.Description = description
}

func (u *UpdateSchemaRequestParameters) SetClientTokenId(clientTokenId string) {
    u.ClientTokenId = clientTokenId
}

func (u *UpdateSchemaRequestParameters) SetContent(content string) {
    u.Content = content
}
