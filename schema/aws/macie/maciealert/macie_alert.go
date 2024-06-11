package maciealert

type MacieAlert struct {
    Summary Summary `json:"summary,omitempty"`
    Trigger Trigger `json:"trigger"`
    Actor string `json:"actor"`
    RiskScore float64 `json:"risk-score"`
    NotificationType string `json:"notification-type"`
    Name string `json:"name"`
    CreatedAt string `json:"created-at"`
    Url string `json:"url"`
    Tags []string `json:"tags"`
    AlertArn string `json:"alert-arn"`
}

func (m *MacieAlert) SetSummary(summary Summary) {
    m.Summary = summary
}

func (m *MacieAlert) SetTrigger(trigger Trigger) {
    m.Trigger = trigger
}

func (m *MacieAlert) SetActor(actor string) {
    m.Actor = actor
}

func (m *MacieAlert) SetRiskScore(riskScore float64) {
    m.RiskScore = riskScore
}

func (m *MacieAlert) SetNotificationType(notificationType string) {
    m.NotificationType = notificationType
}

func (m *MacieAlert) SetName(name string) {
    m.Name = name
}

func (m *MacieAlert) SetCreatedAt(createdAt string) {
    m.CreatedAt = createdAt
}

func (m *MacieAlert) SetUrl(url string) {
    m.Url = url
}

func (m *MacieAlert) SetTags(tags []string) {
    m.Tags = tags
}

func (m *MacieAlert) SetAlertArn(alertArn string) {
    m.AlertArn = alertArn
}
