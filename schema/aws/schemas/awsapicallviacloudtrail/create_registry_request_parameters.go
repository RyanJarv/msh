package awsapicallviacloudtrail

type CreateRegistryRequestParameters struct {
    RegistryName string `json:"registryName"`
    Description string `json:"Description,omitempty"`
    Tags Tags `json:"tags,omitempty"`
}

func (c *CreateRegistryRequestParameters) SetRegistryName(registryName string) {
    c.RegistryName = registryName
}

func (c *CreateRegistryRequestParameters) SetDescription(description string) {
    c.Description = description
}

func (c *CreateRegistryRequestParameters) SetTags(tags Tags) {
    c.Tags = tags
}
