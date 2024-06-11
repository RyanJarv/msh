package devopsgurunewanomalyassociation

type CloudWatchDimension struct {
    Name string `json:"name,omitempty"`
    Value string `json:"value,omitempty"`
}

func (c *CloudWatchDimension) SetName(name string) {
    c.Name = name
}

func (c *CloudWatchDimension) SetValue(value string) {
    c.Value = value
}
