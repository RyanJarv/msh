package awsxrayinsightupdate

type Event struct {
    ClientRequestImpactStatistics ClientRequestImpactStatistics `json:"ClientRequestImpactStatistics"`
    RootCauseServiceRequestImpactStatistics ClientRequestImpactStatistics `json:"RootCauseServiceRequestImpactStatistics"`
    EventTime string `json:"EventTime"`
    Summary string `json:"Summary"`
    TopAnomalousServices []AWSXRayInsightUpdateItem `json:"TopAnomalousServices"`
}

func (e *Event) SetClientRequestImpactStatistics(clientRequestImpactStatistics ClientRequestImpactStatistics) {
    e.ClientRequestImpactStatistics = clientRequestImpactStatistics
}

func (e *Event) SetRootCauseServiceRequestImpactStatistics(rootCauseServiceRequestImpactStatistics ClientRequestImpactStatistics) {
    e.RootCauseServiceRequestImpactStatistics = rootCauseServiceRequestImpactStatistics
}

func (e *Event) SetEventTime(eventTime string) {
    e.EventTime = eventTime
}

func (e *Event) SetSummary(summary string) {
    e.Summary = summary
}

func (e *Event) SetTopAnomalousServices(topAnomalousServices []AWSXRayInsightUpdateItem) {
    e.TopAnomalousServices = topAnomalousServices
}
