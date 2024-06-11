package guarddutyfinding

type ContainerDetails struct {
    Id string `json:"id,omitempty"`
    Image string `json:"image,omitempty"`
    Name string `json:"name,omitempty"`
}

func (c *ContainerDetails) SetId(id string) {
    c.Id = id
}

func (c *ContainerDetails) SetImage(image string) {
    c.Image = image
}

func (c *ContainerDetails) SetName(name string) {
    c.Name = name
}
