package contactlenspostcallrulesmatched

type ContactLensPostCallRulesMatched struct {
    ActionName string `json:"actionName"`
    AgentArn string `json:"agentArn"`
    ContactArn string `json:"contactArn"`
    InstanceArn string `json:"instanceArn"`
    QueueArn string `json:"queueArn"`
    RuleName string `json:"ruleName"`
    Version string `json:"version"`
}

func (c *ContactLensPostCallRulesMatched) SetActionName(actionName string) {
    c.ActionName = actionName
}

func (c *ContactLensPostCallRulesMatched) SetAgentArn(agentArn string) {
    c.AgentArn = agentArn
}

func (c *ContactLensPostCallRulesMatched) SetContactArn(contactArn string) {
    c.ContactArn = contactArn
}

func (c *ContactLensPostCallRulesMatched) SetInstanceArn(instanceArn string) {
    c.InstanceArn = instanceArn
}

func (c *ContactLensPostCallRulesMatched) SetQueueArn(queueArn string) {
    c.QueueArn = queueArn
}

func (c *ContactLensPostCallRulesMatched) SetRuleName(ruleName string) {
    c.RuleName = ruleName
}

func (c *ContactLensPostCallRulesMatched) SetVersion(version string) {
    c.Version = version
}
