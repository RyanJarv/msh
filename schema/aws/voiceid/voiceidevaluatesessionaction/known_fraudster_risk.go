package voiceidevaluatesessionaction

type KnownFraudsterRisk struct {
    GeneratedFraudsterId interface{} `json:"generatedFraudsterId,omitempty"`
    RiskScore float64 `json:"riskScore,omitempty"`
}

func (k *KnownFraudsterRisk) SetGeneratedFraudsterId(generatedFraudsterId interface{}) {
    k.GeneratedFraudsterId = generatedFraudsterId
}

func (k *KnownFraudsterRisk) SetRiskScore(riskScore float64) {
    k.RiskScore = riskScore
}
