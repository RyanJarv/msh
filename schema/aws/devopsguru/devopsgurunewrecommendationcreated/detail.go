package devopsgurunewrecommendationcreated

type Detail struct {
    ResourceCollection ResourceCollection `json:"resourceCollection,omitempty"`
    AccountId string `json:"accountId"`
    InsightName string `json:"insightName,omitempty"`
    InsightDescription string `json:"insightDescription,omitempty"`
    InsightId string `json:"insightId"`
    InsightType string `json:"insightType,omitempty"`
    InsightUrl string `json:"insightUrl,omitempty"`
    MessageType string `json:"messageType"`
    Recommendations []Recommendation `json:"recommendations"`
    Region string `json:"region"`
    StartTime string `json:"startTime,omitempty"`
    StartTimeISO string `json:"startTimeISO,omitempty"`
}

func (d *Detail) SetResourceCollection(resourceCollection ResourceCollection) {
    d.ResourceCollection = resourceCollection
}

func (d *Detail) SetAccountId(accountId string) {
    d.AccountId = accountId
}

func (d *Detail) SetInsightName(insightName string) {
    d.InsightName = insightName
}

func (d *Detail) SetInsightDescription(insightDescription string) {
    d.InsightDescription = insightDescription
}

func (d *Detail) SetInsightId(insightId string) {
    d.InsightId = insightId
}

func (d *Detail) SetInsightType(insightType string) {
    d.InsightType = insightType
}

func (d *Detail) SetInsightUrl(insightUrl string) {
    d.InsightUrl = insightUrl
}

func (d *Detail) SetMessageType(messageType string) {
    d.MessageType = messageType
}

func (d *Detail) SetRecommendations(recommendations []Recommendation) {
    d.Recommendations = recommendations
}

func (d *Detail) SetRegion(region string) {
    d.Region = region
}

func (d *Detail) SetStartTime(startTime string) {
    d.StartTime = startTime
}

func (d *Detail) SetStartTimeISO(startTimeISO string) {
    d.StartTimeISO = startTimeISO
}
