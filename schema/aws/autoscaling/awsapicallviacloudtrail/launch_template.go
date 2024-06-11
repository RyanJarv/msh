package awsapicallviacloudtrail

type LaunchTemplate struct {
    LaunchTemplateName string `json:"launchTemplateName,omitempty"`
}

func (l *LaunchTemplate) SetLaunchTemplateName(launchTemplateName string) {
    l.LaunchTemplateName = launchTemplateName
}
