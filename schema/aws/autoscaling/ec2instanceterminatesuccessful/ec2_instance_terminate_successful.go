package ec2instanceterminatesuccessful

import (
    "time"
)


type EC2InstanceTerminateSuccessful struct {
    Details Details `json:"Details"`
    Description string `json:"Description"`
    EndTime time.Time `json:"EndTime"`
    RequestId string `json:"RequestId"`
    ActivityId string `json:"ActivityId"`
    Cause string `json:"Cause"`
    AutoScalingGroupName string `json:"AutoScalingGroupName"`
    StartTime time.Time `json:"StartTime"`
    EC2InstanceId string `json:"EC2InstanceId"`
    StatusCode string `json:"StatusCode"`
    StatusMessage string `json:"StatusMessage"`
    Origin string `json:"Origin"`
    Destination string `json:"Destination"`
}

func (e *EC2InstanceTerminateSuccessful) SetDetails(details Details) {
    e.Details = details
}

func (e *EC2InstanceTerminateSuccessful) SetDescription(description string) {
    e.Description = description
}

func (e *EC2InstanceTerminateSuccessful) SetEndTime(endTime time.Time) {
    e.EndTime = endTime
}

func (e *EC2InstanceTerminateSuccessful) SetRequestId(requestId string) {
    e.RequestId = requestId
}

func (e *EC2InstanceTerminateSuccessful) SetActivityId(activityId string) {
    e.ActivityId = activityId
}

func (e *EC2InstanceTerminateSuccessful) SetCause(cause string) {
    e.Cause = cause
}

func (e *EC2InstanceTerminateSuccessful) SetAutoScalingGroupName(autoScalingGroupName string) {
    e.AutoScalingGroupName = autoScalingGroupName
}

func (e *EC2InstanceTerminateSuccessful) SetStartTime(startTime time.Time) {
    e.StartTime = startTime
}

func (e *EC2InstanceTerminateSuccessful) SetEC2InstanceId(eC2InstanceId string) {
    e.EC2InstanceId = eC2InstanceId
}

func (e *EC2InstanceTerminateSuccessful) SetStatusCode(statusCode string) {
    e.StatusCode = statusCode
}

func (e *EC2InstanceTerminateSuccessful) SetStatusMessage(statusMessage string) {
    e.StatusMessage = statusMessage
}

func (e *EC2InstanceTerminateSuccessful) SetOrigin(origin string) {
    e.Origin = origin
}

func (e *EC2InstanceTerminateSuccessful) SetDestination(destination string) {
    e.Destination = destination
}
