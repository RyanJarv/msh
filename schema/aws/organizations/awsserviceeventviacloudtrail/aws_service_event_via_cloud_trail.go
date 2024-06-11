package awsserviceeventviacloudtrail

import (
    "time"
)


type AWSServiceEventViaCloudTrail struct {
    ServiceEventDetails ServiceEventDetails `json:"serviceEventDetails"`
    UserIdentity UserIdentity `json:"userIdentity"`
    EventID string `json:"eventID"`
    AwsRegion string `json:"awsRegion"`
    EventVersion string `json:"eventVersion"`
    SourceIPAddress string `json:"sourceIPAddress"`
    EventTime time.Time `json:"eventTime"`
    EventSource string `json:"eventSource"`
    EventName string `json:"eventName"`
    UserAgent string `json:"userAgent"`
    ReadOnly bool `json:"readOnly"`
    EventType string `json:"eventType"`
}

func (a *AWSServiceEventViaCloudTrail) SetServiceEventDetails(serviceEventDetails ServiceEventDetails) {
    a.ServiceEventDetails = serviceEventDetails
}

func (a *AWSServiceEventViaCloudTrail) SetUserIdentity(userIdentity UserIdentity) {
    a.UserIdentity = userIdentity
}

func (a *AWSServiceEventViaCloudTrail) SetEventID(eventID string) {
    a.EventID = eventID
}

func (a *AWSServiceEventViaCloudTrail) SetAwsRegion(awsRegion string) {
    a.AwsRegion = awsRegion
}

func (a *AWSServiceEventViaCloudTrail) SetEventVersion(eventVersion string) {
    a.EventVersion = eventVersion
}

func (a *AWSServiceEventViaCloudTrail) SetSourceIPAddress(sourceIPAddress string) {
    a.SourceIPAddress = sourceIPAddress
}

func (a *AWSServiceEventViaCloudTrail) SetEventTime(eventTime time.Time) {
    a.EventTime = eventTime
}

func (a *AWSServiceEventViaCloudTrail) SetEventSource(eventSource string) {
    a.EventSource = eventSource
}

func (a *AWSServiceEventViaCloudTrail) SetEventName(eventName string) {
    a.EventName = eventName
}

func (a *AWSServiceEventViaCloudTrail) SetUserAgent(userAgent string) {
    a.UserAgent = userAgent
}

func (a *AWSServiceEventViaCloudTrail) SetReadOnly(readOnly bool) {
    a.ReadOnly = readOnly
}

func (a *AWSServiceEventViaCloudTrail) SetEventType(eventType string) {
    a.EventType = eventType
}
