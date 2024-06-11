package awsapicallviacloudtrail

type ResponseElements struct {
    Build Build `json:"build,omitempty"`
    Project Project `json:"project,omitempty"`
    Webhook Webhook `json:"webhook,omitempty"`
    Message string `json:"message,omitempty"`
    WebhookDeletedStatus string `json:"webhookDeletedStatus,omitempty"`
}

func (r *ResponseElements) SetBuild(build Build) {
    r.Build = build
}

func (r *ResponseElements) SetProject(project Project) {
    r.Project = project
}

func (r *ResponseElements) SetWebhook(webhook Webhook) {
    r.Webhook = webhook
}

func (r *ResponseElements) SetMessage(message string) {
    r.Message = message
}

func (r *ResponseElements) SetWebhookDeletedStatus(webhookDeletedStatus string) {
    r.WebhookDeletedStatus = webhookDeletedStatus
}
