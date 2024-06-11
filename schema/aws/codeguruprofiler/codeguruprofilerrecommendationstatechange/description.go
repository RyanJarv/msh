package codeguruprofilerrecommendationstatechange

type Description struct {
    Value string `json:"value"`
}

func (d *Description) SetValue(value string) {
    d.Value = value
}
