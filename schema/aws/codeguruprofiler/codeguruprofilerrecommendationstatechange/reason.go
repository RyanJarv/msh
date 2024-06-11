package codeguruprofilerrecommendationstatechange

type Reason struct {
    Value string `json:"value"`
}

func (r *Reason) SetValue(value string) {
    r.Value = value
}
