package awsapicallviacloudtrail

type ContainerOverridesItem struct {
    Name string `json:"name"`
    Value string `json:"value"`
}

func (c *ContainerOverridesItem) SetName(name string) {
    c.Name = name
}

func (c *ContainerOverridesItem) SetValue(value string) {
    c.Value = value
}
