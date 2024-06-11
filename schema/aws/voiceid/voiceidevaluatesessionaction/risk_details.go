package voiceidevaluatesessionaction

type RiskDetails struct {
    KnownFraudsterRisk KnownFraudsterRisk `json:"knownFraudsterRisk,omitempty"`
    VoiceSpoofingRisk VoiceSpoofingRisk `json:"voiceSpoofingRisk,omitempty"`
}

func (r *RiskDetails) SetKnownFraudsterRisk(knownFraudsterRisk KnownFraudsterRisk) {
    r.KnownFraudsterRisk = knownFraudsterRisk
}

func (r *RiskDetails) SetVoiceSpoofingRisk(voiceSpoofingRisk VoiceSpoofingRisk) {
    r.VoiceSpoofingRisk = voiceSpoofingRisk
}
