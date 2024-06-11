package codedeployinstancestatechangenotification

type CodeDeployInstanceStateChangeNotification struct {
    InstanceId string `json:"instanceId"`
    Application string `json:"application"`
    DeploymentId string `json:"deploymentId"`
    DeploymentGroup string `json:"deploymentGroup"`
    InstanceGroupId string `json:"instanceGroupId"`
    State string `json:"state"`
    Region string `json:"region"`
}

func (c *CodeDeployInstanceStateChangeNotification) SetInstanceId(instanceId string) {
    c.InstanceId = instanceId
}

func (c *CodeDeployInstanceStateChangeNotification) SetApplication(application string) {
    c.Application = application
}

func (c *CodeDeployInstanceStateChangeNotification) SetDeploymentId(deploymentId string) {
    c.DeploymentId = deploymentId
}

func (c *CodeDeployInstanceStateChangeNotification) SetDeploymentGroup(deploymentGroup string) {
    c.DeploymentGroup = deploymentGroup
}

func (c *CodeDeployInstanceStateChangeNotification) SetInstanceGroupId(instanceGroupId string) {
    c.InstanceGroupId = instanceGroupId
}

func (c *CodeDeployInstanceStateChangeNotification) SetState(state string) {
    c.State = state
}

func (c *CodeDeployInstanceStateChangeNotification) SetRegion(region string) {
    c.Region = region
}
