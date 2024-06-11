package voiceidevaluatesessionaction

import (
    "time"
)


type AuthenticationResult struct {
    Configuration Configuration_Authentication `json:"configuration,omitempty"`
    AudioAggregationEndedAt time.Time `json:"audioAggregationEndedAt,omitempty"`
    AudioAggregationStartedAt time.Time `json:"audioAggregationStartedAt,omitempty"`
    AuthenticationResultId string `json:"authenticationResultId,omitempty"`
    Decision string `json:"decision,omitempty"`
    Score float64 `json:"score,omitempty"`
}

func (a *AuthenticationResult) SetConfiguration(configuration Configuration_Authentication) {
    a.Configuration = configuration
}

func (a *AuthenticationResult) SetAudioAggregationEndedAt(audioAggregationEndedAt time.Time) {
    a.AudioAggregationEndedAt = audioAggregationEndedAt
}

func (a *AuthenticationResult) SetAudioAggregationStartedAt(audioAggregationStartedAt time.Time) {
    a.AudioAggregationStartedAt = audioAggregationStartedAt
}

func (a *AuthenticationResult) SetAuthenticationResultId(authenticationResultId string) {
    a.AuthenticationResultId = authenticationResultId
}

func (a *AuthenticationResult) SetDecision(decision string) {
    a.Decision = decision
}

func (a *AuthenticationResult) SetScore(score float64) {
    a.Score = score
}
