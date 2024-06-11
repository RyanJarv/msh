package awsapicallviacloudtrail

type LaunchTemplate_1 struct {
    LaunchTemplateSpecification LaunchTemplateSpecification `json:"launchTemplateSpecification,omitempty"`
    Overrides []interface{} `json:"overrides,omitempty"`
}

func (l *LaunchTemplate_1) SetLaunchTemplateSpecification(launchTemplateSpecification LaunchTemplateSpecification) {
    l.LaunchTemplateSpecification = launchTemplateSpecification
}

func (l *LaunchTemplate_1) SetOverrides(overrides []interface{}) {
    l.Overrides = overrides
}
