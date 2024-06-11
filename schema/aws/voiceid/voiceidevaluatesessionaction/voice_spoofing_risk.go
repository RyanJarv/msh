package voiceidevaluatesessionaction

type VoiceSpoofingRisk struct {
    RiskScore float64 `json:"riskScore,omitempty"`
}

func (v *VoiceSpoofingRisk) SetRiskScore(riskScore float64) {
    v.RiskScore = riskScore
}
