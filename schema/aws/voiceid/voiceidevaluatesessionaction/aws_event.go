package voiceidevaluatesessionaction

import (
    "time"
)


type AWSEvent struct {
    Detail VoiceIdEvaluateSessionAction `json:"detail,omitempty"`
    Account string `json:"account,omitempty"`
    DetailType string `json:"detail-type,omitempty"`
    Id string `json:"id,omitempty"`
    Region string `json:"region,omitempty"`
    Resources []string `json:"resources,omitempty"`
    Source string `json:"source,omitempty"`
    Time time.Time `json:"time,omitempty"`
    Version string `json:"version,omitempty"`
}

func (a *AWSEvent) SetDetail(detail VoiceIdEvaluateSessionAction) {
    a.Detail = detail
}

func (a *AWSEvent) SetAccount(account string) {
    a.Account = account
}

func (a *AWSEvent) SetDetailType(detailType string) {
    a.DetailType = detailType
}

func (a *AWSEvent) SetId(id string) {
    a.Id = id
}

func (a *AWSEvent) SetRegion(region string) {
    a.Region = region
}

func (a *AWSEvent) SetResources(resources []string) {
    a.Resources = resources
}

func (a *AWSEvent) SetSource(source string) {
    a.Source = source
}

func (a *AWSEvent) SetTime(time time.Time) {
    a.Time = time
}

func (a *AWSEvent) SetVersion(version string) {
    a.Version = version
}
