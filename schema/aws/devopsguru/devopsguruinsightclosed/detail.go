package devopsguruinsightclosed

type Detail struct {
    ResourceCollection ResourceCollection `json:"resourceCollection,omitempty"`
    AccountId string `json:"accountId,omitempty"`
    Anomalies []Anomaly `json:"anomalies,omitempty"`
    InsightDescription string `json:"insightDescription,omitempty"`
    InsightName string `json:"insightName,omitempty"`
    InsightId string `json:"insightId,omitempty"`
    InsightSeverity string `json:"insightSeverity,omitempty"`
    InsightType string `json:"insightType,omitempty"`
    InsightUrl string `json:"insightUrl,omitempty"`
    MessageType string `json:"messageType,omitempty"`
    Region string `json:"region,omitempty"`
    StartTime string `json:"startTime,omitempty"`
    StartTimeISO string `json:"startTimeISO,omitempty"`
    EndTime string `json:"endTime,omitempty"`
    EndTimeISO string `json:"endTimeISO,omitempty"`
}

func (d *Detail) SetResourceCollection(resourceCollection ResourceCollection) {
    d.ResourceCollection = resourceCollection
}

func (d *Detail) SetAccountId(accountId string) {
    d.AccountId = accountId
}

func (d *Detail) SetAnomalies(anomalies []Anomaly) {
    d.Anomalies = anomalies
}

func (d *Detail) SetInsightDescription(insightDescription string) {
    d.InsightDescription = insightDescription
}

func (d *Detail) SetInsightName(insightName string) {
    d.InsightName = insightName
}

func (d *Detail) SetInsightId(insightId string) {
    d.InsightId = insightId
}

func (d *Detail) SetInsightSeverity(insightSeverity string) {
    d.InsightSeverity = insightSeverity
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

func (d *Detail) SetRegion(region string) {
    d.Region = region
}

func (d *Detail) SetStartTime(startTime string) {
    d.StartTime = startTime
}

func (d *Detail) SetStartTimeISO(startTimeISO string) {
    d.StartTimeISO = startTimeISO
}

func (d *Detail) SetEndTime(endTime string) {
    d.EndTime = endTime
}

func (d *Detail) SetEndTimeISO(endTimeISO string) {
    d.EndTimeISO = endTimeISO
}
