package awsxrayinsightupdate

type AWSXRayInsightUpdate struct {
    ClientRequestImpactStatistics ClientRequestImpactStatistics `json:"ClientRequestImpactStatistics"`
    Event Event `json:"Event"`
    RootCauseServiceId RootCauseServiceId `json:"RootCauseServiceId"`
    RootCauseServiceRequestImpactStatistics ClientRequestImpactStatistics `json:"RootCauseServiceRequestImpactStatistics"`
    Categories []string `json:"Categories"`
    EndTime interface{} `json:"EndTime"`
    GroupName string `json:"GroupName"`
    InsightId string `json:"InsightId"`
    StartTime float64 `json:"StartTime"`
    State string `json:"State"`
    Summary string `json:"Summary"`
    TopAnomalousServices []AWSXRayInsightUpdateItem `json:"TopAnomalousServices"`
}

func (a *AWSXRayInsightUpdate) SetClientRequestImpactStatistics(clientRequestImpactStatistics ClientRequestImpactStatistics) {
    a.ClientRequestImpactStatistics = clientRequestImpactStatistics
}

func (a *AWSXRayInsightUpdate) SetEvent(event Event) {
    a.Event = event
}

func (a *AWSXRayInsightUpdate) SetRootCauseServiceId(rootCauseServiceId RootCauseServiceId) {
    a.RootCauseServiceId = rootCauseServiceId
}

func (a *AWSXRayInsightUpdate) SetRootCauseServiceRequestImpactStatistics(rootCauseServiceRequestImpactStatistics ClientRequestImpactStatistics) {
    a.RootCauseServiceRequestImpactStatistics = rootCauseServiceRequestImpactStatistics
}

func (a *AWSXRayInsightUpdate) SetCategories(categories []string) {
    a.Categories = categories
}

func (a *AWSXRayInsightUpdate) SetEndTime(endTime interface{}) {
    a.EndTime = endTime
}

func (a *AWSXRayInsightUpdate) SetGroupName(groupName string) {
    a.GroupName = groupName
}

func (a *AWSXRayInsightUpdate) SetInsightId(insightId string) {
    a.InsightId = insightId
}

func (a *AWSXRayInsightUpdate) SetStartTime(startTime float64) {
    a.StartTime = startTime
}

func (a *AWSXRayInsightUpdate) SetState(state string) {
    a.State = state
}

func (a *AWSXRayInsightUpdate) SetSummary(summary string) {
    a.Summary = summary
}

func (a *AWSXRayInsightUpdate) SetTopAnomalousServices(topAnomalousServices []AWSXRayInsightUpdateItem) {
    a.TopAnomalousServices = topAnomalousServices
}
