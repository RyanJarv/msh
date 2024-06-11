package awshealthabuseevent

type AWSHealthAbuseEvent struct {
    EventArn string `json:"eventArn,omitempty"`
    EventTypeCode string `json:"eventTypeCode,omitempty"`
    Service string `json:"service,omitempty"`
    EventDescription []EventDescription `json:"eventDescription,omitempty"`
    StartTime string `json:"startTime,omitempty"`
    EndTime string `json:"endTime,omitempty"`
    EventTypeCategory string `json:"eventTypeCategory,omitempty"`
    AffectedEntities []AffectedEntity `json:"affectedEntities,omitempty"`
    CommunicationId string `json:"communicationId,omitempty"`
    EventRegion string `json:"eventRegion,omitempty"`
    EventScopeCode string `json:"eventScopeCode,omitempty"`
    LastUpdatedTime string `json:"lastUpdatedTime,omitempty"`
    StatusCode string `json:"statusCode,omitempty"`
    EventMetadata interface{} `json:"eventMetadata,omitempty"`
}

func (a *AWSHealthAbuseEvent) SetEventArn(eventArn string) {
    a.EventArn = eventArn
}

func (a *AWSHealthAbuseEvent) SetEventTypeCode(eventTypeCode string) {
    a.EventTypeCode = eventTypeCode
}

func (a *AWSHealthAbuseEvent) SetService(service string) {
    a.Service = service
}

func (a *AWSHealthAbuseEvent) SetEventDescription(eventDescription []EventDescription) {
    a.EventDescription = eventDescription
}

func (a *AWSHealthAbuseEvent) SetStartTime(startTime string) {
    a.StartTime = startTime
}

func (a *AWSHealthAbuseEvent) SetEndTime(endTime string) {
    a.EndTime = endTime
}

func (a *AWSHealthAbuseEvent) SetEventTypeCategory(eventTypeCategory string) {
    a.EventTypeCategory = eventTypeCategory
}

func (a *AWSHealthAbuseEvent) SetAffectedEntities(affectedEntities []AffectedEntity) {
    a.AffectedEntities = affectedEntities
}

func (a *AWSHealthAbuseEvent) SetCommunicationId(communicationId string) {
    a.CommunicationId = communicationId
}

func (a *AWSHealthAbuseEvent) SetEventRegion(eventRegion string) {
    a.EventRegion = eventRegion
}

func (a *AWSHealthAbuseEvent) SetEventScopeCode(eventScopeCode string) {
    a.EventScopeCode = eventScopeCode
}

func (a *AWSHealthAbuseEvent) SetLastUpdatedTime(lastUpdatedTime string) {
    a.LastUpdatedTime = lastUpdatedTime
}

func (a *AWSHealthAbuseEvent) SetStatusCode(statusCode string) {
    a.StatusCode = statusCode
}

func (a *AWSHealthAbuseEvent) SetEventMetadata(eventMetadata interface{}) {
    a.EventMetadata = eventMetadata
}
