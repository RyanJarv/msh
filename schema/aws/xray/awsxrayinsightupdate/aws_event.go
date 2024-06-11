package awsxrayinsightupdate

import (
    "time"
)


type AWSEvent struct {
    Detail AWSXRayInsightUpdate `json:"detail"`
    Account string `json:"account"`
    DetailType string `json:"detail-type"`
    Id string `json:"id"`
    Region string `json:"region"`
    Source string `json:"source"`
    Time time.Time `json:"time"`
    Version string `json:"version"`
}

func (a *AWSEvent) SetDetail(detail AWSXRayInsightUpdate) {
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

func (a *AWSEvent) SetSource(source string) {
    a.Source = source
}

func (a *AWSEvent) SetTime(time time.Time) {
    a.Time = time
}

func (a *AWSEvent) SetVersion(version string) {
    a.Version = version
}
