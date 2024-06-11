package voiceidupdatesessionaction

type Session struct {
    AuthenticationConfiguration AuthenticationConfiguration `json:"authenticationConfiguration,omitempty"`
    FraudDetectionConfiguration FraudDetectionConfiguration `json:"fraudDetectionConfiguration,omitempty"`
    GeneratedSpeakerId string `json:"generatedSpeakerId,omitempty"`
    SessionId string `json:"sessionId,omitempty"`
    SessionName string `json:"sessionName,omitempty"`
}

func (s *Session) SetAuthenticationConfiguration(authenticationConfiguration AuthenticationConfiguration) {
    s.AuthenticationConfiguration = authenticationConfiguration
}

func (s *Session) SetFraudDetectionConfiguration(fraudDetectionConfiguration FraudDetectionConfiguration) {
    s.FraudDetectionConfiguration = fraudDetectionConfiguration
}

func (s *Session) SetGeneratedSpeakerId(generatedSpeakerId string) {
    s.GeneratedSpeakerId = generatedSpeakerId
}

func (s *Session) SetSessionId(sessionId string) {
    s.SessionId = sessionId
}

func (s *Session) SetSessionName(sessionName string) {
    s.SessionName = sessionName
}
