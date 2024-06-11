package voiceidstartsessionaction

import (
    "time"
)


type AuthenticationAudioProgress struct {
    AudioAggregationEndedAt time.Time `json:"audioAggregationEndedAt,omitempty"`
    AudioAggregationStartedAt time.Time `json:"audioAggregationStartedAt,omitempty"`
}

func (a *AuthenticationAudioProgress) SetAudioAggregationEndedAt(audioAggregationEndedAt time.Time) {
    a.AudioAggregationEndedAt = audioAggregationEndedAt
}

func (a *AuthenticationAudioProgress) SetAudioAggregationStartedAt(audioAggregationStartedAt time.Time) {
    a.AudioAggregationStartedAt = audioAggregationStartedAt
}
