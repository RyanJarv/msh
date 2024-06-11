package awsserviceeventviacloudtrail

import (
    "time"
)


type AWSServiceEventViaCloudTrail struct {
    ServiceEventDetails ServiceEventDetails `json:"serviceEventDetails"`
    UserIdentity UserIdentity `json:"userIdentity"`
    AwsRegion string `json:"awsRegion"`
    EventID string `json:"eventID"`
    EventName string `json:"eventName"`
    EventSource string `json:"eventSource"`
    EventTime time.Time `json:"eventTime"`
    EventType string `json:"eventType"`
    EventVersion string `json:"eventVersion"`
    ReadOnly bool `json:"readOnly"`
    SourceIPAddress string `json:"sourceIPAddress"`
    UserAgent string `json:"userAgent"`
}

func (a *AWSServiceEventViaCloudTrail) SetServiceEventDetails(serviceEventDetails ServiceEventDetails) {
    a.ServiceEventDetails = serviceEventDetails
}

func (a *AWSServiceEventViaCloudTrail) SetUserIdentity(userIdentity UserIdentity) {
    a.UserIdentity = userIdentity
}

func (a *AWSServiceEventViaCloudTrail) SetAwsRegion(awsRegion string) {
    a.AwsRegion = awsRegion
}

func (a *AWSServiceEventViaCloudTrail) SetEventID(eventID string) {
    a.EventID = eventID
}

func (a *AWSServiceEventViaCloudTrail) SetEventName(eventName string) {
    a.EventName = eventName
}

func (a *AWSServiceEventViaCloudTrail) SetEventSource(eventSource string) {
    a.EventSource = eventSource
}

func (a *AWSServiceEventViaCloudTrail) SetEventTime(eventTime time.Time) {
    a.EventTime = eventTime
}

func (a *AWSServiceEventViaCloudTrail) SetEventType(eventType string) {
    a.EventType = eventType
}

func (a *AWSServiceEventViaCloudTrail) SetEventVersion(eventVersion string) {
    a.EventVersion = eventVersion
}

func (a *AWSServiceEventViaCloudTrail) SetReadOnly(readOnly bool) {
    a.ReadOnly = readOnly
}

func (a *AWSServiceEventViaCloudTrail) SetSourceIPAddress(sourceIPAddress string) {
    a.SourceIPAddress = sourceIPAddress
}

func (a *AWSServiceEventViaCloudTrail) SetUserAgent(userAgent string) {
    a.UserAgent = userAgent
}
