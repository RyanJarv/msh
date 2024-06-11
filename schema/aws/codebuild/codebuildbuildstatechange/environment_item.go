package codebuildbuildstatechange

type EnvironmentItem struct {
    Name string `json:"name,omitempty"`
    EnvironmentItemType string `json:"type,omitempty"`
    Value string `json:"value,omitempty"`
}

func (e *EnvironmentItem) SetName(name string) {
    e.Name = name
}

func (e *EnvironmentItem) SetEnvironmentItemType(environmentItemType string) {
    e.EnvironmentItemType = environmentItemType
}

func (e *EnvironmentItem) SetValue(value string) {
    e.Value = value
}
