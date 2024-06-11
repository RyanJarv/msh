package awsapicallviacloudtrail

type MixedInstancesPolicy struct {
    LaunchTemplate LaunchTemplate_1 `json:"launchTemplate,omitempty"`
    InstancesDistribution InstancesDistribution `json:"instancesDistribution,omitempty"`
}

func (m *MixedInstancesPolicy) SetLaunchTemplate(launchTemplate LaunchTemplate_1) {
    m.LaunchTemplate = launchTemplate
}

func (m *MixedInstancesPolicy) SetInstancesDistribution(instancesDistribution InstancesDistribution) {
    m.InstancesDistribution = instancesDistribution
}
