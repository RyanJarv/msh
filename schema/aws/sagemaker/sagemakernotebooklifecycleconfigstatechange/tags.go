package sagemakernotebooklifecycleconfigstatechange

type Tags struct {
    Key string `json:"Key"`
    Value string `json:"Value"`
}

func (t *Tags) SetKey(key string) {
    t.Key = key
}

func (t *Tags) SetValue(value string) {
    t.Value = value
}
