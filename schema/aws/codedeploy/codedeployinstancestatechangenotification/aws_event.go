package codedeployinstancestatechangenotification

import (
    "time"
)


type AWSEvent struct {
    Detail CodeDeployInstanceStateChangeNotification `json:"detail"`
    DetailType string `json:"detail-type"`
    Resources []string `json:"resources"`
    Source string `json:"source"`
    Time time.Time `json:"time"`
    Id string `json:"id"`
    Region string `json:"region"`
    Version string `json:"version"`
    Account string `json:"account"`
}

func (a *AWSEvent) SetDetail(detail CodeDeployInstanceStateChangeNotification) {
    a.Detail = detail
}

func (a *AWSEvent) SetDetailType(detailType string) {
    a.DetailType = detailType
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

func (a *AWSEvent) SetId(id string) {
    a.Id = id
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
