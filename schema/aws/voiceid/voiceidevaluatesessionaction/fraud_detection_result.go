package voiceidevaluatesessionaction

import (
    "time"
)


type FraudDetectionResult struct {
    Configuration Configuration_Fraud `json:"configuration,omitempty"`
    RiskDetails RiskDetails `json:"riskDetails,omitempty"`
    AudioAggregationEndedAt time.Time `json:"audioAggregationEndedAt,omitempty"`
    AudioAggregationStartedAt time.Time `json:"audioAggregationStartedAt,omitempty"`
    Decision string `json:"decision,omitempty"`
    FraudDetectionResultId string `json:"fraudDetectionResultId,omitempty"`
    Reasons []string `json:"reasons,omitempty"`
}

func (f *FraudDetectionResult) SetConfiguration(configuration Configuration_Fraud) {
    f.Configuration = configuration
}

func (f *FraudDetectionResult) SetRiskDetails(riskDetails RiskDetails) {
    f.RiskDetails = riskDetails
}

func (f *FraudDetectionResult) SetAudioAggregationEndedAt(audioAggregationEndedAt time.Time) {
    f.AudioAggregationEndedAt = audioAggregationEndedAt
}

func (f *FraudDetectionResult) SetAudioAggregationStartedAt(audioAggregationStartedAt time.Time) {
    f.AudioAggregationStartedAt = audioAggregationStartedAt
}

func (f *FraudDetectionResult) SetDecision(decision string) {
    f.Decision = decision
}

func (f *FraudDetectionResult) SetFraudDetectionResultId(fraudDetectionResultId string) {
    f.FraudDetectionResultId = fraudDetectionResultId
}

func (f *FraudDetectionResult) SetReasons(reasons []string) {
    f.Reasons = reasons
}
