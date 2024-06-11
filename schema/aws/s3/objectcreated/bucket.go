package objectcreated

type Bucket struct {
    Name string `json:"name"`
}

func (b *Bucket) SetName(name string) {
    b.Name = name
}
