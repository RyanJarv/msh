package voiceidstartsessionaction

type Session struct {
    AuthenticationAudioProgress AuthenticationAudioProgress `json:"authenticationAudioProgress,omitempty"`
    AuthenticationConfiguration AuthenticationConfiguration `json:"authenticationConfiguration,omitempty"`
    EnrollmentAudioProgress EnrollmentAudioProgress `json:"enrollmentAudioProgress,omitempty"`
    FraudDetectionAudioProgress AuthenticationAudioProgress `json:"fraudDetectionAudioProgress,omitempty"`
    FraudDetectionConfiguration FraudDetectionConfiguration `json:"fraudDetectionConfiguration,omitempty"`
    StreamingConfiguration StreamingConfiguration `json:"streamingConfiguration,omitempty"`
    GeneratedSpeakerId string `json:"generatedSpeakerId,omitempty"`
    SessionId string `json:"sessionId,omitempty"`
    SessionName string `json:"sessionName,omitempty"`
}

func (s *Session) SetAuthenticationAudioProgress(authenticationAudioProgress AuthenticationAudioProgress) {
    s.AuthenticationAudioProgress = authenticationAudioProgress
}

func (s *Session) SetAuthenticationConfiguration(authenticationConfiguration AuthenticationConfiguration) {
    s.AuthenticationConfiguration = authenticationConfiguration
}

func (s *Session) SetEnrollmentAudioProgress(enrollmentAudioProgress EnrollmentAudioProgress) {
    s.EnrollmentAudioProgress = enrollmentAudioProgress
}

func (s *Session) SetFraudDetectionAudioProgress(fraudDetectionAudioProgress AuthenticationAudioProgress) {
    s.FraudDetectionAudioProgress = fraudDetectionAudioProgress
}

func (s *Session) SetFraudDetectionConfiguration(fraudDetectionConfiguration FraudDetectionConfiguration) {
    s.FraudDetectionConfiguration = fraudDetectionConfiguration
}

func (s *Session) SetStreamingConfiguration(streamingConfiguration StreamingConfiguration) {
    s.StreamingConfiguration = streamingConfiguration
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
