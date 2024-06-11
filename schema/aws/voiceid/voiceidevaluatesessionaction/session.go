package voiceidevaluatesessionaction

type Session struct {
    AuthenticationResult AuthenticationResult `json:"authenticationResult,omitempty"`
    FraudDetectionResult FraudDetectionResult `json:"fraudDetectionResult,omitempty"`
    GeneratedSpeakerId string `json:"generatedSpeakerId,omitempty"`
    SessionId string `json:"sessionId,omitempty"`
    SessionName string `json:"sessionName,omitempty"`
    StreamingStatus string `json:"streamingStatus,omitempty"`
}

func (s *Session) SetAuthenticationResult(authenticationResult AuthenticationResult) {
    s.AuthenticationResult = authenticationResult
}

func (s *Session) SetFraudDetectionResult(fraudDetectionResult FraudDetectionResult) {
    s.FraudDetectionResult = fraudDetectionResult
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

func (s *Session) SetStreamingStatus(streamingStatus string) {
    s.StreamingStatus = streamingStatus
}
