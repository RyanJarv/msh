package codeconnectcontact

type AgentInfo struct {
    AgentArn string `json:"agentArn"`
}

func (a *AgentInfo) SetAgentArn(agentArn string) {
    a.AgentArn = agentArn
}
