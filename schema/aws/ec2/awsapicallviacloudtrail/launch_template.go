package awsapicallviacloudtrail

type LaunchTemplate struct {
    LaunchTemplateId string `json:"launchTemplateId,omitempty"`
    Version string `json:"version,omitempty"`
}

func (l *LaunchTemplate) SetLaunchTemplateId(launchTemplateId string) {
    l.LaunchTemplateId = launchTemplateId
}

func (l *LaunchTemplate) SetVersion(version string) {
    l.Version = version
}
