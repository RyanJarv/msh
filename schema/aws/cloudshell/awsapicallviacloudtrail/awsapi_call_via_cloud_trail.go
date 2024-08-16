package awsapicallviacloudtrail

import (
    "time"
)


type AWSAPICallViaCloudTrail struct {
    RequestParameters RequestParameters `json:"requestParameters"`
    ResponseElements ResponseElements `json:"responseElements"`
    UserIdentity UserIdentity `json:"userIdentity"`
    AwsRegion string `json:"awsRegion"`
    EventCategory string `json:"eventCategory"`
    EventID string `json:"eventID"`
    EventName string `json:"eventName"`
    EventSource string `json:"eventSource"`
    EventTime time.Time `json:"eventTime"`
    EventType string `json:"eventType"`
    EventVersion string `json:"eventVersion"`
    ManagementEvent bool `json:"managementEvent"`
    ReadOnly bool `json:"readOnly"`
    RecipientAccountId string `json:"recipientAccountId"`
    RequestID string `json:"requestID"`
    SourceIPAddress string `json:"sourceIPAddress"`
    UserAgent string `json:"userAgent"`
}

func (a *AWSAPICallViaCloudTrail) SetRequestParameters(requestParameters RequestParameters) {
    a.RequestParameters = requestParameters
}

func (a *AWSAPICallViaCloudTrail) SetResponseElements(responseElements ResponseElements) {
    a.ResponseElements = responseElements
}

func (a *AWSAPICallViaCloudTrail) SetUserIdentity(userIdentity UserIdentity) {
    a.UserIdentity = userIdentity
}

func (a *AWSAPICallViaCloudTrail) SetAwsRegion(awsRegion string) {
    a.AwsRegion = awsRegion
}

func (a *AWSAPICallViaCloudTrail) SetEventCategory(eventCategory string) {
    a.EventCategory = eventCategory
}

func (a *AWSAPICallViaCloudTrail) SetEventID(eventID string) {
    a.EventID = eventID
}

func (a *AWSAPICallViaCloudTrail) SetEventName(eventName string) {
    a.EventName = eventName
}

func (a *AWSAPICallViaCloudTrail) SetEventSource(eventSource string) {
    a.EventSource = eventSource
}

func (a *AWSAPICallViaCloudTrail) SetEventTime(eventTime time.Time) {
    a.EventTime = eventTime
}

func (a *AWSAPICallViaCloudTrail) SetEventType(eventType string) {
    a.EventType = eventType
}

func (a *AWSAPICallViaCloudTrail) SetEventVersion(eventVersion string) {
    a.EventVersion = eventVersion
}

func (a *AWSAPICallViaCloudTrail) SetManagementEvent(managementEvent bool) {
    a.ManagementEvent = managementEvent
}

func (a *AWSAPICallViaCloudTrail) SetReadOnly(readOnly bool) {
    a.ReadOnly = readOnly
}

func (a *AWSAPICallViaCloudTrail) SetRecipientAccountId(recipientAccountId string) {
    a.RecipientAccountId = recipientAccountId
}

func (a *AWSAPICallViaCloudTrail) SetRequestID(requestID string) {
    a.RequestID = requestID
}

func (a *AWSAPICallViaCloudTrail) SetSourceIPAddress(sourceIPAddress string) {
    a.SourceIPAddress = sourceIPAddress
}

func (a *AWSAPICallViaCloudTrail) SetUserAgent(userAgent string) {
    a.UserAgent = userAgent
}