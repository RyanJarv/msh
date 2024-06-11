package opsworksdeploymentstatechange

type OpsWorksDeploymentStateChange struct {
    Duration float64 `json:"duration"`
    StackId string `json:"stack-id"`
    InstanceIds []string `json:"instance-ids"`
    DeploymentId string `json:"deployment-id"`
    Command string `json:"command"`
    Status string `json:"status"`
}

func (o *OpsWorksDeploymentStateChange) SetDuration(duration float64) {
    o.Duration = duration
}

func (o *OpsWorksDeploymentStateChange) SetStackId(stackId string) {
    o.StackId = stackId
}

func (o *OpsWorksDeploymentStateChange) SetInstanceIds(instanceIds []string) {
    o.InstanceIds = instanceIds
}

func (o *OpsWorksDeploymentStateChange) SetDeploymentId(deploymentId string) {
    o.DeploymentId = deploymentId
}

func (o *OpsWorksDeploymentStateChange) SetCommand(command string) {
    o.Command = command
}

func (o *OpsWorksDeploymentStateChange) SetStatus(status string) {
    o.Status = status
}
