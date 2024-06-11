package awsconsolesigninviacloudtrail

import (
    "time"
)


type AWSConsoleSignInViaCloudTrail struct {
    ResponseElements ResponseElements `json:"responseElements"`
    UserIdentity UserIdentity `json:"userIdentity"`
    AdditionalEventData AdditionalEventData `json:"additionalEventData"`
    EventID string `json:"eventID"`
    AwsRegion string `json:"awsRegion"`
    EventVersion string `json:"eventVersion"`
    SourceIPAddress string `json:"sourceIPAddress"`
    EventSource string `json:"eventSource"`
    RequestParameters interface{} `json:"requestParameters"`
    UserAgent string `json:"userAgent"`
    EventType string `json:"eventType"`
    EventTime time.Time `json:"eventTime"`
    EventName string `json:"eventName"`
}

func (a *AWSConsoleSignInViaCloudTrail) SetResponseElements(responseElements ResponseElements) {
    a.ResponseElements = responseElements
}

func (a *AWSConsoleSignInViaCloudTrail) SetUserIdentity(userIdentity UserIdentity) {
    a.UserIdentity = userIdentity
}

func (a *AWSConsoleSignInViaCloudTrail) SetAdditionalEventData(additionalEventData AdditionalEventData) {
    a.AdditionalEventData = additionalEventData
}

func (a *AWSConsoleSignInViaCloudTrail) SetEventID(eventID string) {
    a.EventID = eventID
}

func (a *AWSConsoleSignInViaCloudTrail) SetAwsRegion(awsRegion string) {
    a.AwsRegion = awsRegion
}

func (a *AWSConsoleSignInViaCloudTrail) SetEventVersion(eventVersion string) {
    a.EventVersion = eventVersion
}

func (a *AWSConsoleSignInViaCloudTrail) SetSourceIPAddress(sourceIPAddress string) {
    a.SourceIPAddress = sourceIPAddress
}

func (a *AWSConsoleSignInViaCloudTrail) SetEventSource(eventSource string) {
    a.EventSource = eventSource
}

func (a *AWSConsoleSignInViaCloudTrail) SetRequestParameters(requestParameters interface{}) {
    a.RequestParameters = requestParameters
}

func (a *AWSConsoleSignInViaCloudTrail) SetUserAgent(userAgent string) {
    a.UserAgent = userAgent
}

func (a *AWSConsoleSignInViaCloudTrail) SetEventType(eventType string) {
    a.EventType = eventType
}

func (a *AWSConsoleSignInViaCloudTrail) SetEventTime(eventTime time.Time) {
    a.EventTime = eventTime
}

func (a *AWSConsoleSignInViaCloudTrail) SetEventName(eventName string) {
    a.EventName = eventName
}
