package awsapicallviacloudtrail

type CreateLaunchTemplateRequest struct {
    LaunchTemplateData LaunchTemplateData `json:"LaunchTemplateData,omitempty"`
    LaunchTemplateName string `json:"LaunchTemplateName,omitempty"`
}

func (c *CreateLaunchTemplateRequest) SetLaunchTemplateData(launchTemplateData LaunchTemplateData) {
    c.LaunchTemplateData = launchTemplateData
}

func (c *CreateLaunchTemplateRequest) SetLaunchTemplateName(launchTemplateName string) {
    c.LaunchTemplateName = launchTemplateName
}
