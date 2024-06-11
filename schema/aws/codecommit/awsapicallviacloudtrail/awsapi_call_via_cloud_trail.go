package awsapicallviacloudtrail

import (
    "time"
)


type AWSAPICallViaCloudTrail struct {
    AdditionalEventData AdditionalEventData `json:"additionalEventData,omitempty"`
    RequestParameters RequestParameters `json:"requestParameters"`
    ResponseElements ResponseElements `json:"responseElements"`
    UserIdentity UserIdentity `json:"userIdentity"`
    ApiVersion string `json:"apiVersion,omitempty"`
    AwsRegion string `json:"awsRegion"`
    ErrorCode string `json:"errorCode,omitempty"`
    ErrorMessage string `json:"errorMessage,omitempty"`
    EventID string `json:"eventID"`
    EventName string `json:"eventName"`
    EventSource string `json:"eventSource"`
    EventTime time.Time `json:"eventTime"`
    EventType string `json:"eventType"`
    EventVersion string `json:"eventVersion"`
    ReadOnly bool `json:"readOnly"`
    RequestID string `json:"requestID"`
    Resources []AWSAPICallViaCloudTrailItem `json:"resources,omitempty"`
    SourceIPAddress string `json:"sourceIPAddress"`
    UserAgent string `json:"userAgent"`
}

func (a *AWSAPICallViaCloudTrail) SetAdditionalEventData(additionalEventData AdditionalEventData) {
    a.AdditionalEventData = additionalEventData
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

func (a *AWSAPICallViaCloudTrail) SetApiVersion(apiVersion string) {
    a.ApiVersion = apiVersion
}

func (a *AWSAPICallViaCloudTrail) SetAwsRegion(awsRegion string) {
    a.AwsRegion = awsRegion
}

func (a *AWSAPICallViaCloudTrail) SetErrorCode(errorCode string) {
    a.ErrorCode = errorCode
}

func (a *AWSAPICallViaCloudTrail) SetErrorMessage(errorMessage string) {
    a.ErrorMessage = errorMessage
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

func (a *AWSAPICallViaCloudTrail) SetReadOnly(readOnly bool) {
    a.ReadOnly = readOnly
}

func (a *AWSAPICallViaCloudTrail) SetRequestID(requestID string) {
    a.RequestID = requestID
}

func (a *AWSAPICallViaCloudTrail) SetResources(resources []AWSAPICallViaCloudTrailItem) {
    a.Resources = resources
}

func (a *AWSAPICallViaCloudTrail) SetSourceIPAddress(sourceIPAddress string) {
    a.SourceIPAddress = sourceIPAddress
}

func (a *AWSAPICallViaCloudTrail) SetUserAgent(userAgent string) {
    a.UserAgent = userAgent
}
