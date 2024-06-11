package awsapicallviacloudtrail

import (
    "time"
)


type AWSAPICallViaCloudTrail struct {
    RequestParameters RequestParameters `json:"requestParameters"`
    UserIdentity UserIdentity `json:"userIdentity"`
    EventID string `json:"eventID"`
    AwsRegion string `json:"awsRegion"`
    EventVersion string `json:"eventVersion"`
    ResponseElements interface{} `json:"responseElements"`
    SourceIPAddress string `json:"sourceIPAddress"`
    EventSource string `json:"eventSource"`
    Resources []AWSAPICallViaCloudTrailItem `json:"resources"`
    UserAgent string `json:"userAgent"`
    ReadOnly bool `json:"readOnly"`
    EventType string `json:"eventType"`
    RequestID string `json:"requestID"`
    EventTime time.Time `json:"eventTime"`
    EventName string `json:"eventName"`
}

func (a *AWSAPICallViaCloudTrail) SetRequestParameters(requestParameters RequestParameters) {
    a.RequestParameters = requestParameters
}

func (a *AWSAPICallViaCloudTrail) SetUserIdentity(userIdentity UserIdentity) {
    a.UserIdentity = userIdentity
}

func (a *AWSAPICallViaCloudTrail) SetEventID(eventID string) {
    a.EventID = eventID
}

func (a *AWSAPICallViaCloudTrail) SetAwsRegion(awsRegion string) {
    a.AwsRegion = awsRegion
}

func (a *AWSAPICallViaCloudTrail) SetEventVersion(eventVersion string) {
    a.EventVersion = eventVersion
}

func (a *AWSAPICallViaCloudTrail) SetResponseElements(responseElements interface{}) {
    a.ResponseElements = responseElements
}

func (a *AWSAPICallViaCloudTrail) SetSourceIPAddress(sourceIPAddress string) {
    a.SourceIPAddress = sourceIPAddress
}

func (a *AWSAPICallViaCloudTrail) SetEventSource(eventSource string) {
    a.EventSource = eventSource
}

func (a *AWSAPICallViaCloudTrail) SetResources(resources []AWSAPICallViaCloudTrailItem) {
    a.Resources = resources
}

func (a *AWSAPICallViaCloudTrail) SetUserAgent(userAgent string) {
    a.UserAgent = userAgent
}

func (a *AWSAPICallViaCloudTrail) SetReadOnly(readOnly bool) {
    a.ReadOnly = readOnly
}

func (a *AWSAPICallViaCloudTrail) SetEventType(eventType string) {
    a.EventType = eventType
}

func (a *AWSAPICallViaCloudTrail) SetRequestID(requestID string) {
    a.RequestID = requestID
}

func (a *AWSAPICallViaCloudTrail) SetEventTime(eventTime time.Time) {
    a.EventTime = eventTime
}

func (a *AWSAPICallViaCloudTrail) SetEventName(eventName string) {
    a.EventName = eventName
}
