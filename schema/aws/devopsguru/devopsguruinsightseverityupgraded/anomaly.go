package devopsguruinsightseverityupgraded

type Anomaly struct {
    AnomalyResources []AnomalyResource `json:"anomalyResources,omitempty"`
    AssociatedResourceArns []string `json:"associatedResourceArns,omitempty"`
    Description string `json:"description,omitempty"`
    EndTime string `json:"endTime,omitempty"`
    Id string `json:"id,omitempty"`
    SourceDetails []SourceDetail `json:"sourceDetails,omitempty"`
    StartTime string `json:"startTime,omitempty"`
    StartTimeISO string `json:"startTimeISO,omitempty"`
    OpenTime string `json:"openTime,omitempty"`
    OpenTimeISO string `json:"openTimeISO,omitempty"`
    EarliestImpactTime string `json:"earliestImpactTime,omitempty"`
    EarliestImpactTimeISO string `json:"earliestImpactTimeISO,omitempty"`
    LatestImpactTime string `json:"latestImpactTime,omitempty"`
    LatestImpactTimeISO string `json:"latestImpactTimeISO,omitempty"`
    Limit float64 `json:"limit,omitempty"`
    AnomalyType string `json:"type,omitempty"`
    AnomalySeverity string `json:"anomalySeverity,omitempty"`
    SourceMetadata AnomalySourceMetadata `json:"sourceMetadata,omitempty"`
}

func (a *Anomaly) SetAnomalyResources(anomalyResources []AnomalyResource) {
    a.AnomalyResources = anomalyResources
}

func (a *Anomaly) SetAssociatedResourceArns(associatedResourceArns []string) {
    a.AssociatedResourceArns = associatedResourceArns
}

func (a *Anomaly) SetDescription(description string) {
    a.Description = description
}

func (a *Anomaly) SetEndTime(endTime string) {
    a.EndTime = endTime
}

func (a *Anomaly) SetId(id string) {
    a.Id = id
}

func (a *Anomaly) SetSourceDetails(sourceDetails []SourceDetail) {
    a.SourceDetails = sourceDetails
}

func (a *Anomaly) SetStartTime(startTime string) {
    a.StartTime = startTime
}

func (a *Anomaly) SetStartTimeISO(startTimeISO string) {
    a.StartTimeISO = startTimeISO
}

func (a *Anomaly) SetOpenTime(openTime string) {
    a.OpenTime = openTime
}

func (a *Anomaly) SetOpenTimeISO(openTimeISO string) {
    a.OpenTimeISO = openTimeISO
}

func (a *Anomaly) SetEarliestImpactTime(earliestImpactTime string) {
    a.EarliestImpactTime = earliestImpactTime
}

func (a *Anomaly) SetEarliestImpactTimeISO(earliestImpactTimeISO string) {
    a.EarliestImpactTimeISO = earliestImpactTimeISO
}

func (a *Anomaly) SetLatestImpactTime(latestImpactTime string) {
    a.LatestImpactTime = latestImpactTime
}

func (a *Anomaly) SetLatestImpactTimeISO(latestImpactTimeISO string) {
    a.LatestImpactTimeISO = latestImpactTimeISO
}

func (a *Anomaly) SetLimit(limit float64) {
    a.Limit = limit
}

func (a *Anomaly) SetAnomalyType(anomalyType string) {
    a.AnomalyType = anomalyType
}

func (a *Anomaly) SetAnomalySeverity(anomalySeverity string) {
    a.AnomalySeverity = anomalySeverity
}

func (a *Anomaly) SetSourceMetadata(sourceMetadata AnomalySourceMetadata) {
    a.SourceMetadata = sourceMetadata
}
