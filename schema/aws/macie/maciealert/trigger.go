package maciealert

type Trigger struct {
    Features Features `json:"features,omitempty"`
    RuleArn string `json:"rule-arn,omitempty"`
    AlertType string `json:"alert-type"`
    Description string `json:"description,omitempty"`
    CreatedAt string `json:"created-at,omitempty"`
    Risk float64 `json:"risk,omitempty"`
}

func (t *Trigger) SetFeatures(features Features) {
    t.Features = features
}

func (t *Trigger) SetRuleArn(ruleArn string) {
    t.RuleArn = ruleArn
}

func (t *Trigger) SetAlertType(alertType string) {
    t.AlertType = alertType
}

func (t *Trigger) SetDescription(description string) {
    t.Description = description
}

func (t *Trigger) SetCreatedAt(createdAt string) {
    t.CreatedAt = createdAt
}

func (t *Trigger) SetRisk(risk float64) {
    t.Risk = risk
}
