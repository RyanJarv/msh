package codeguruprofilerrecommendationstatechange

type Name struct {
    Value string `json:"value"`
}

func (n *Name) SetValue(value string) {
    n.Value = value
}
