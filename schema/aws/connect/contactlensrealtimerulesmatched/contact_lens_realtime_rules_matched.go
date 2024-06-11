package contactlensrealtimerulesmatched

type ContactLensRealtimeRulesMatched struct {
    ActionName string `json:"actionName"`
    AgentArn string `json:"agentArn"`
    ContactArn string `json:"contactArn"`
    InstanceArn string `json:"instanceArn"`
    QueueArn string `json:"queueArn"`
    RuleName string `json:"ruleName"`
    Version string `json:"version"`
}

func (c *ContactLensRealtimeRulesMatched) SetActionName(actionName string) {
    c.ActionName = actionName
}

func (c *ContactLensRealtimeRulesMatched) SetAgentArn(agentArn string) {
    c.AgentArn = agentArn
}

func (c *ContactLensRealtimeRulesMatched) SetContactArn(contactArn string) {
    c.ContactArn = contactArn
}

func (c *ContactLensRealtimeRulesMatched) SetInstanceArn(instanceArn string) {
    c.InstanceArn = instanceArn
}

func (c *ContactLensRealtimeRulesMatched) SetQueueArn(queueArn string) {
    c.QueueArn = queueArn
}

func (c *ContactLensRealtimeRulesMatched) SetRuleName(ruleName string) {
    c.RuleName = ruleName
}

func (c *ContactLensRealtimeRulesMatched) SetVersion(version string) {
    c.Version = version
}
