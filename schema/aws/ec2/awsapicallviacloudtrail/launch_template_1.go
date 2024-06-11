package awsapicallviacloudtrail

import (
    "time"
)


type LaunchTemplate_1 struct {
    CreateTime time.Time `json:"createTime,omitempty"`
    CreatedBy string `json:"createdBy,omitempty"`
    LaunchTemplateId string `json:"launchTemplateId,omitempty"`
    LatestVersionNumber float64 `json:"latestVersionNumber,omitempty"`
    DefaultVersionNumber float64 `json:"defaultVersionNumber,omitempty"`
    LaunchTemplateName string `json:"launchTemplateName,omitempty"`
}

func (l *LaunchTemplate_1) SetCreateTime(createTime time.Time) {
    l.CreateTime = createTime
}

func (l *LaunchTemplate_1) SetCreatedBy(createdBy string) {
    l.CreatedBy = createdBy
}

func (l *LaunchTemplate_1) SetLaunchTemplateId(launchTemplateId string) {
    l.LaunchTemplateId = launchTemplateId
}

func (l *LaunchTemplate_1) SetLatestVersionNumber(latestVersionNumber float64) {
    l.LatestVersionNumber = latestVersionNumber
}

func (l *LaunchTemplate_1) SetDefaultVersionNumber(defaultVersionNumber float64) {
    l.DefaultVersionNumber = defaultVersionNumber
}

func (l *LaunchTemplate_1) SetLaunchTemplateName(launchTemplateName string) {
    l.LaunchTemplateName = launchTemplateName
}
