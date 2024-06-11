package chimevoiceconnectorstreamingstatus

type AWSEvent struct {
    Detail ChimeVoiceConnectorStreamingStatus `json:"detail"`
    Account string `json:"account"`
    DetailType string `json:"detail-type"`
    Id string `json:"id"`
    Region string `json:"region"`
    Resources []interface{} `json:"resources"`
    Source string `json:"source"`
    Time string `json:"time"`
    Version string `json:"version"`
}

func (a *AWSEvent) SetDetail(detail ChimeVoiceConnectorStreamingStatus) {
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

func (a *AWSEvent) SetResources(resources []interface{}) {
    a.Resources = resources
}

func (a *AWSEvent) SetSource(source string) {
    a.Source = source
}

func (a *AWSEvent) SetTime(time string) {
    a.Time = time
}

func (a *AWSEvent) SetVersion(version string) {
    a.Version = version
}
