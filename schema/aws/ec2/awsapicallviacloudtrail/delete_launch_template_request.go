package awsapicallviacloudtrail

type DeleteLaunchTemplateRequest struct {
    LaunchTemplateName string `json:"LaunchTemplateName,omitempty"`
}

func (d *DeleteLaunchTemplateRequest) SetLaunchTemplateName(launchTemplateName string) {
    d.LaunchTemplateName = launchTemplateName
}
