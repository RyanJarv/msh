package awshealthabuseevent

import (
    "time"
)


type AWSEvent struct {
    Detail AWSHealthAbuseEvent `json:"detail,omitempty"`
    DetailType string `json:"detail-type,omitempty"`
    Resources []string `json:"resources,omitempty"`
    Id string `json:"id,omitempty"`
    Source string `json:"source,omitempty"`
    Time time.Time `json:"time,omitempty"`
    Region string `json:"region,omitempty"`
    Version string `json:"version,omitempty"`
    Account string `json:"account,omitempty"`
}

func (a *AWSEvent) SetDetail(detail AWSHealthAbuseEvent) {
    a.Detail = detail
}

func (a *AWSEvent) SetDetailType(detailType string) {
    a.DetailType = detailType
}

func (a *AWSEvent) SetResources(resources []string) {
    a.Resources = resources
}

func (a *AWSEvent) SetId(id string) {
    a.Id = id
}

func (a *AWSEvent) SetSource(source string) {
    a.Source = source
}

func (a *AWSEvent) SetTime(time time.Time) {
    a.Time = time
}

func (a *AWSEvent) SetRegion(region string) {
    a.Region = region
}

func (a *AWSEvent) SetVersion(version string) {
    a.Version = version
}

func (a *AWSEvent) SetAccount(account string) {
    a.Account = account
}
