package voiceidstartsessionaction

type StreamingConfiguration struct {
    AuthenticationMinimumSpeechInSeconds float64 `json:"authenticationMinimumSpeechInSeconds,omitempty"`
}

func (s *StreamingConfiguration) SetAuthenticationMinimumSpeechInSeconds(authenticationMinimumSpeechInSeconds float64) {
    s.AuthenticationMinimumSpeechInSeconds = authenticationMinimumSpeechInSeconds
}
