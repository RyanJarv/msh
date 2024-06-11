package codeguruprofilerrecommendationstatechange

type Title struct {
    Value string `json:"value"`
}

func (t *Title) SetValue(value string) {
    t.Value = value
}
