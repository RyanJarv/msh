package codedeploydeploymentstatechangenotification

type CodeDeployDeploymentStateChangeNotification struct {
    Application string `json:"application"`
    DeploymentId string `json:"deploymentId"`
    DeploymentGroup string `json:"deploymentGroup"`
    InstanceGroupId string `json:"instanceGroupId"`
    State string `json:"state"`
    Region string `json:"region"`
}

func (c *CodeDeployDeploymentStateChangeNotification) SetApplication(application string) {
    c.Application = application
}

func (c *CodeDeployDeploymentStateChangeNotification) SetDeploymentId(deploymentId string) {
    c.DeploymentId = deploymentId
}

func (c *CodeDeployDeploymentStateChangeNotification) SetDeploymentGroup(deploymentGroup string) {
    c.DeploymentGroup = deploymentGroup
}

func (c *CodeDeployDeploymentStateChangeNotification) SetInstanceGroupId(instanceGroupId string) {
    c.InstanceGroupId = instanceGroupId
}

func (c *CodeDeployDeploymentStateChangeNotification) SetState(state string) {
    c.State = state
}

func (c *CodeDeployDeploymentStateChangeNotification) SetRegion(region string) {
    c.Region = region
}
