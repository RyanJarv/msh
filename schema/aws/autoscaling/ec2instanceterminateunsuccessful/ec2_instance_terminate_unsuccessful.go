package ec2instanceterminateunsuccessful

import (
    "time"
)


type EC2InstanceTerminateUnsuccessful struct {
    Details Details `json:"Details"`
    Description string `json:"Description"`
    RequestId string `json:"RequestId"`
    EndTime time.Time `json:"EndTime"`
    AutoScalingGroupName string `json:"AutoScalingGroupName"`
    ActivityId string `json:"ActivityId"`
    Cause string `json:"Cause"`
    StartTime time.Time `json:"StartTime"`
    EC2InstanceId string `json:"EC2InstanceId"`
    StatusCode string `json:"StatusCode"`
    StatusMessage string `json:"StatusMessage"`
    Origin string `json:"Origin"`
    Destination string `json:"Destination"`
}

func (e *EC2InstanceTerminateUnsuccessful) SetDetails(details Details) {
    e.Details = details
}

func (e *EC2InstanceTerminateUnsuccessful) SetDescription(description string) {
    e.Description = description
}

func (e *EC2InstanceTerminateUnsuccessful) SetRequestId(requestId string) {
    e.RequestId = requestId
}

func (e *EC2InstanceTerminateUnsuccessful) SetEndTime(endTime time.Time) {
    e.EndTime = endTime
}

func (e *EC2InstanceTerminateUnsuccessful) SetAutoScalingGroupName(autoScalingGroupName string) {
    e.AutoScalingGroupName = autoScalingGroupName
}

func (e *EC2InstanceTerminateUnsuccessful) SetActivityId(activityId string) {
    e.ActivityId = activityId
}

func (e *EC2InstanceTerminateUnsuccessful) SetCause(cause string) {
    e.Cause = cause
}

func (e *EC2InstanceTerminateUnsuccessful) SetStartTime(startTime time.Time) {
    e.StartTime = startTime
}

func (e *EC2InstanceTerminateUnsuccessful) SetEC2InstanceId(eC2InstanceId string) {
    e.EC2InstanceId = eC2InstanceId
}

func (e *EC2InstanceTerminateUnsuccessful) SetStatusCode(statusCode string) {
    e.StatusCode = statusCode
}

func (e *EC2InstanceTerminateUnsuccessful) SetStatusMessage(statusMessage string) {
    e.StatusMessage = statusMessage
}

func (e *EC2InstanceTerminateUnsuccessful) SetOrigin(origin string) {
    e.Origin = origin
}

func (e *EC2InstanceTerminateUnsuccessful) SetDestination(destination string) {
    e.Destination = destination
}
