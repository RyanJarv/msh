package awsapicallviacloudtrail

import (
    "time"
)


type AWSAPICallViaCloudTrail struct {
    RequestParameters RequestParameters `json:"requestParameters"`
    TlsDetails TlsDetails `json:"tlsDetails,omitempty"`
    UserIdentity UserIdentity `json:"userIdentity"`
    AdditionalEventData AdditionalEventData `json:"additionalEventData"`
    EventID string `json:"eventID"`
    AwsRegion string `json:"awsRegion"`
    EventCategory string `json:"eventCategory,omitempty"`
    ResponseElements interface{} `json:"responseElements"`
    SourceIPAddress string `json:"sourceIPAddress"`
    EventName string `json:"eventName"`
    EventSource string `json:"eventSource"`
    EventTime time.Time `json:"eventTime"`
    EventVersion string `json:"eventVersion"`
    ManagementEvent bool `json:"managementEvent,omitempty"`
    ErrorMessage string `json:"errorMessage,omitempty"`
    Resources []AWSAPICallViaCloudTrailItem `json:"resources"`
    ErrorCode string `json:"errorCode,omitempty"`
    UserAgent string `json:"userAgent"`
    ReadOnly bool `json:"readOnly"`
    EventType string `json:"eventType"`
    VpcEndpointId string `json:"vpcEndpointId,omitempty"`
    RequestID string `json:"requestID"`
    RecipientAccountId string `json:"recipientAccountId"`
}

func (a *AWSAPICallViaCloudTrail) SetRequestParameters(requestParameters RequestParameters) {
    a.RequestParameters = requestParameters
}

func (a *AWSAPICallViaCloudTrail) SetTlsDetails(tlsDetails TlsDetails) {
    a.TlsDetails = tlsDetails
}

func (a *AWSAPICallViaCloudTrail) SetUserIdentity(userIdentity UserIdentity) {
    a.UserIdentity = userIdentity
}

func (a *AWSAPICallViaCloudTrail) SetAdditionalEventData(additionalEventData AdditionalEventData) {
    a.AdditionalEventData = additionalEventData
}

func (a *AWSAPICallViaCloudTrail) SetEventID(eventID string) {
    a.EventID = eventID
}

func (a *AWSAPICallViaCloudTrail) SetAwsRegion(awsRegion string) {
    a.AwsRegion = awsRegion
}

func (a *AWSAPICallViaCloudTrail) SetEventCategory(eventCategory string) {
    a.EventCategory = eventCategory
}

func (a *AWSAPICallViaCloudTrail) SetResponseElements(responseElements interface{}) {
    a.ResponseElements = responseElements
}

func (a *AWSAPICallViaCloudTrail) SetSourceIPAddress(sourceIPAddress string) {
    a.SourceIPAddress = sourceIPAddress
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

func (a *AWSAPICallViaCloudTrail) SetEventVersion(eventVersion string) {
    a.EventVersion = eventVersion
}

func (a *AWSAPICallViaCloudTrail) SetManagementEvent(managementEvent bool) {
    a.ManagementEvent = managementEvent
}

func (a *AWSAPICallViaCloudTrail) SetErrorMessage(errorMessage string) {
    a.ErrorMessage = errorMessage
}

func (a *AWSAPICallViaCloudTrail) SetResources(resources []AWSAPICallViaCloudTrailItem) {
    a.Resources = resources
}

func (a *AWSAPICallViaCloudTrail) SetErrorCode(errorCode string) {
    a.ErrorCode = errorCode
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

func (a *AWSAPICallViaCloudTrail) SetVpcEndpointId(vpcEndpointId string) {
    a.VpcEndpointId = vpcEndpointId
}

func (a *AWSAPICallViaCloudTrail) SetRequestID(requestID string) {
    a.RequestID = requestID
}

func (a *AWSAPICallViaCloudTrail) SetRecipientAccountId(recipientAccountId string) {
    a.RecipientAccountId = recipientAccountId
}
