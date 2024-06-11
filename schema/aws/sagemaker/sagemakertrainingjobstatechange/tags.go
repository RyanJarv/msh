package sagemakertrainingjobstatechange

type Tags struct {
    Key string `json:"Key,omitempty"`
    Value string `json:"Value,omitempty"`
}

func (t *Tags) SetKey(key string) {
    t.Key = key
}

func (t *Tags) SetValue(value string) {
    t.Value = value
}
