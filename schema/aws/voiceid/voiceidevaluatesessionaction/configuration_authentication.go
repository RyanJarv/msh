package voiceidevaluatesessionaction

type Configuration_Authentication struct {
    AcceptanceThreshold float64 `json:"acceptanceThreshold,omitempty"`
}

func (c *Configuration_Authentication) SetAcceptanceThreshold(acceptanceThreshold float64) {
    c.AcceptanceThreshold = acceptanceThreshold
}
