package awsapicallviacloudtrail

type LaunchTemplateConfigs struct {
    LaunchTemplateSpecification LaunchTemplateSpecification `json:"LaunchTemplateSpecification,omitempty"`
    Overrides []interface{} `json:"Overrides,omitempty"`
    Tag float64 `json:"tag,omitempty"`
}

func (l *LaunchTemplateConfigs) SetLaunchTemplateSpecification(launchTemplateSpecification LaunchTemplateSpecification) {
    l.LaunchTemplateSpecification = launchTemplateSpecification
}

func (l *LaunchTemplateConfigs) SetOverrides(overrides []interface{}) {
    l.Overrides = overrides
}

func (l *LaunchTemplateConfigs) SetTag(tag float64) {
    l.Tag = tag
}
