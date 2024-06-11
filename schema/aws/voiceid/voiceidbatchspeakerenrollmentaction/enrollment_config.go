package voiceidbatchspeakerenrollmentaction

type EnrollmentConfig struct {
    FraudDetectionConfig FraudDetectionConfig `json:"fraudDetectionConfig,omitempty"`
    ExistingEnrollmentAction string `json:"existingEnrollmentAction,omitempty"`
}

func (e *EnrollmentConfig) SetFraudDetectionConfig(fraudDetectionConfig FraudDetectionConfig) {
    e.FraudDetectionConfig = fraudDetectionConfig
}

func (e *EnrollmentConfig) SetExistingEnrollmentAction(existingEnrollmentAction string) {
    e.ExistingEnrollmentAction = existingEnrollmentAction
}
