package awsapicallviacloudtrail

type CreateDiscovererRequestParameters struct {
    CrossAccount bool `json:"CrossAccount,omitempty"`
    Description string `json:"Description,omitempty"`
    SourceArn string `json:"SourceArn"`
    Tags Tags `json:"tags,omitempty"`
}

func (c *CreateDiscovererRequestParameters) SetCrossAccount(crossAccount bool) {
    c.CrossAccount = crossAccount
}

func (c *CreateDiscovererRequestParameters) SetDescription(description string) {
    c.Description = description
}

func (c *CreateDiscovererRequestParameters) SetSourceArn(sourceArn string) {
    c.SourceArn = sourceArn
}

func (c *CreateDiscovererRequestParameters) SetTags(tags Tags) {
    c.Tags = tags
}
