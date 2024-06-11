package awsapicallviacloudtrail

type LaunchTemplateSpecification struct {
    Version string `json:"version,omitempty"`
    LaunchTemplateName string `json:"launchTemplateName,omitempty"`
}

func (l *LaunchTemplateSpecification) SetVersion(version string) {
    l.Version = version
}

func (l *LaunchTemplateSpecification) SetLaunchTemplateName(launchTemplateName string) {
    l.LaunchTemplateName = launchTemplateName
}
