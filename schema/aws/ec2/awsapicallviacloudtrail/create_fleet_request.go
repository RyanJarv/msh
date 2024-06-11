package awsapicallviacloudtrail

type CreateFleetRequest struct {
    TargetCapacitySpecification TargetCapacitySpecification `json:"TargetCapacitySpecification,omitempty"`
    OnDemandOptions OnDemandOptions `json:"OnDemandOptions,omitempty"`
    SpotOptions SpotOptions `json:"SpotOptions,omitempty"`
    ExistingInstances ExistingInstances `json:"ExistingInstances,omitempty"`
    LaunchTemplateConfigs LaunchTemplateConfigs `json:"LaunchTemplateConfigs,omitempty"`
    TagSpecification TagSpecification `json:"TagSpecification,omitempty"`
    CreateFleetRequestType string `json:"Type,omitempty"`
    ClientToken string `json:"ClientToken,omitempty"`
}

func (c *CreateFleetRequest) SetTargetCapacitySpecification(targetCapacitySpecification TargetCapacitySpecification) {
    c.TargetCapacitySpecification = targetCapacitySpecification
}

func (c *CreateFleetRequest) SetOnDemandOptions(onDemandOptions OnDemandOptions) {
    c.OnDemandOptions = onDemandOptions
}

func (c *CreateFleetRequest) SetSpotOptions(spotOptions SpotOptions) {
    c.SpotOptions = spotOptions
}

func (c *CreateFleetRequest) SetExistingInstances(existingInstances ExistingInstances) {
    c.ExistingInstances = existingInstances
}

func (c *CreateFleetRequest) SetLaunchTemplateConfigs(launchTemplateConfigs LaunchTemplateConfigs) {
    c.LaunchTemplateConfigs = launchTemplateConfigs
}

func (c *CreateFleetRequest) SetTagSpecification(tagSpecification TagSpecification) {
    c.TagSpecification = tagSpecification
}

func (c *CreateFleetRequest) SetCreateFleetRequestType(createFleetRequestType string) {
    c.CreateFleetRequestType = createFleetRequestType
}

func (c *CreateFleetRequest) SetClientToken(clientToken string) {
    c.ClientToken = clientToken
}
