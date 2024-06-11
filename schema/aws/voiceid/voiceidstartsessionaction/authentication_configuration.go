package voiceidstartsessionaction

type AuthenticationConfiguration struct {
    AcceptanceThreshold float64 `json:"acceptanceThreshold,omitempty"`
}

func (a *AuthenticationConfiguration) SetAcceptanceThreshold(acceptanceThreshold float64) {
    a.AcceptanceThreshold = acceptanceThreshold
}
