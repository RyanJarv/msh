package awshealthevent

type AWSHealthEvent struct {
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

func (a *AWSHealthEvent) SetEventArn(eventArn string) {
    a.EventArn = eventArn
}

func (a *AWSHealthEvent) SetEventTypeCode(eventTypeCode string) {
    a.EventTypeCode = eventTypeCode
}

func (a *AWSHealthEvent) SetService(service string) {
    a.Service = service
}

func (a *AWSHealthEvent) SetEventDescription(eventDescription []EventDescription) {
    a.EventDescription = eventDescription
}

func (a *AWSHealthEvent) SetStartTime(startTime string) {
    a.StartTime = startTime
}

func (a *AWSHealthEvent) SetEndTime(endTime string) {
    a.EndTime = endTime
}

func (a *AWSHealthEvent) SetEventTypeCategory(eventTypeCategory string) {
    a.EventTypeCategory = eventTypeCategory
}

func (a *AWSHealthEvent) SetAffectedEntities(affectedEntities []AffectedEntity) {
    a.AffectedEntities = affectedEntities
}

func (a *AWSHealthEvent) SetCommunicationId(communicationId string) {
    a.CommunicationId = communicationId
}

func (a *AWSHealthEvent) SetEventRegion(eventRegion string) {
    a.EventRegion = eventRegion
}

func (a *AWSHealthEvent) SetEventScopeCode(eventScopeCode string) {
    a.EventScopeCode = eventScopeCode
}

func (a *AWSHealthEvent) SetLastUpdatedTime(lastUpdatedTime string) {
    a.LastUpdatedTime = lastUpdatedTime
}

func (a *AWSHealthEvent) SetStatusCode(statusCode string) {
    a.StatusCode = statusCode
}

func (a *AWSHealthEvent) SetEventMetadata(eventMetadata interface{}) {
    a.EventMetadata = eventMetadata
}
