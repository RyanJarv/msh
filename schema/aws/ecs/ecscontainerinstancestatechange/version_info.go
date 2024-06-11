package ecscontainerinstancestatechange

type VersionInfo struct {
    DockerVersion string `json:"dockerVersion,omitempty"`
    AgentHash string `json:"agentHash,omitempty"`
    AgentVersion string `json:"agentVersion,omitempty"`
}

func (v *VersionInfo) SetDockerVersion(dockerVersion string) {
    v.DockerVersion = dockerVersion
}

func (v *VersionInfo) SetAgentHash(agentHash string) {
    v.AgentHash = agentHash
}

func (v *VersionInfo) SetAgentVersion(agentVersion string) {
    v.AgentVersion = agentVersion
}
