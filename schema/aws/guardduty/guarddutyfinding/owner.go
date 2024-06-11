package guarddutyfinding

type Owner struct {
    Id string `json:"id,omitempty"`
}

func (o *Owner) SetId(id string) {
    o.Id = id
}
