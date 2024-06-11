package awsapicallviacloudtrail

type LaunchTemplateSpecification struct {
    Version float64 `json:"Version,omitempty"`
    LaunchTemplateId string `json:"LaunchTemplateId,omitempty"`
}

func (l *LaunchTemplateSpecification) SetVersion(version float64) {
    l.Version = version
}

func (l *LaunchTemplateSpecification) SetLaunchTemplateId(launchTemplateId string) {
    l.LaunchTemplateId = launchTemplateId
}
