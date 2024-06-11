package batchjobstatechange

type ContainerItem struct {
    Name string `json:"name"`
    Value string `json:"value"`
}

func (c *ContainerItem) SetName(name string) {
    c.Name = name
}

func (c *ContainerItem) SetValue(value string) {
    c.Value = value
}
