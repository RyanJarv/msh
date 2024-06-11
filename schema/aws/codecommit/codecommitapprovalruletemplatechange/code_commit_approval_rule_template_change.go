package codecommitapprovalruletemplatechange

type CodeCommitApprovalRuleTemplateChange struct {
    Repositories Repositories `json:"repositories"`
    ApprovalRuleTemplateContentSha256 string `json:"approvalRuleTemplateContentSha256"`
    ApprovalRuleTemplateId string `json:"approvalRuleTemplateId"`
    ApprovalRuleTemplateName string `json:"approvalRuleTemplateName"`
    CallerUserArn string `json:"callerUserArn"`
    CreationDate string `json:"creationDate"`
    Event string `json:"event"`
    LastModifiedDate string `json:"lastModifiedDate"`
    NotificationBody string `json:"notificationBody"`
}

func (c *CodeCommitApprovalRuleTemplateChange) SetRepositories(repositories Repositories) {
    c.Repositories = repositories
}

func (c *CodeCommitApprovalRuleTemplateChange) SetApprovalRuleTemplateContentSha256(approvalRuleTemplateContentSha256 string) {
    c.ApprovalRuleTemplateContentSha256 = approvalRuleTemplateContentSha256
}

func (c *CodeCommitApprovalRuleTemplateChange) SetApprovalRuleTemplateId(approvalRuleTemplateId string) {
    c.ApprovalRuleTemplateId = approvalRuleTemplateId
}

func (c *CodeCommitApprovalRuleTemplateChange) SetApprovalRuleTemplateName(approvalRuleTemplateName string) {
    c.ApprovalRuleTemplateName = approvalRuleTemplateName
}

func (c *CodeCommitApprovalRuleTemplateChange) SetCallerUserArn(callerUserArn string) {
    c.CallerUserArn = callerUserArn
}

func (c *CodeCommitApprovalRuleTemplateChange) SetCreationDate(creationDate string) {
    c.CreationDate = creationDate
}

func (c *CodeCommitApprovalRuleTemplateChange) SetEvent(event string) {
    c.Event = event
}

func (c *CodeCommitApprovalRuleTemplateChange) SetLastModifiedDate(lastModifiedDate string) {
    c.LastModifiedDate = lastModifiedDate
}

func (c *CodeCommitApprovalRuleTemplateChange) SetNotificationBody(notificationBody string) {
    c.NotificationBody = notificationBody
}
