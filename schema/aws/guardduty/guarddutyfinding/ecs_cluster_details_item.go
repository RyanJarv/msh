package guarddutyfinding

type EcsClusterDetailsItem struct {
    Key string `json:"key"`
    Value string `json:"value"`
}

func (e *EcsClusterDetailsItem) SetKey(key string) {
    e.Key = key
}

func (e *EcsClusterDetailsItem) SetValue(value string) {
    e.Value = value
}
