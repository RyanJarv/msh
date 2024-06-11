package voiceidstartsessionaction

import (
    "time"
)


type EnrollmentAudioProgress struct {
    AudioAggregationEndedAt time.Time `json:"audioAggregationEndedAt,omitempty"`
    AudioAggregationStartedAt time.Time `json:"audioAggregationStartedAt,omitempty"`
    AudioAggregationStatus string `json:"audioAggregationStatus,omitempty"`
}

func (e *EnrollmentAudioProgress) SetAudioAggregationEndedAt(audioAggregationEndedAt time.Time) {
    e.AudioAggregationEndedAt = audioAggregationEndedAt
}

func (e *EnrollmentAudioProgress) SetAudioAggregationStartedAt(audioAggregationStartedAt time.Time) {
    e.AudioAggregationStartedAt = audioAggregationStartedAt
}

func (e *EnrollmentAudioProgress) SetAudioAggregationStatus(audioAggregationStatus string) {
    e.AudioAggregationStatus = audioAggregationStatus
}
