package codepipelineactionexecutionstatechange

type Type struct {
    Owner string `json:"owner"`
    Provider string `json:"provider"`
    Category string `json:"category"`
    Version float64 `json:"version"`
}

func (t *Type) SetOwner(owner string) {
    t.Owner = owner
}

func (t *Type) SetProvider(provider string) {
    t.Provider = provider
}

func (t *Type) SetCategory(category string) {
    t.Category = category
}

func (t *Type) SetVersion(version float64) {
    t.Version = version
}
