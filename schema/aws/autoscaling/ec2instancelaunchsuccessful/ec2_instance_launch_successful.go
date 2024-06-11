package ec2instancelaunchsuccessful

import (
    "time"
)


type EC2InstanceLaunchSuccessful struct {
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

func (e *EC2InstanceLaunchSuccessful) SetDetails(details Details) {
    e.Details = details
}

func (e *EC2InstanceLaunchSuccessful) SetDescription(description string) {
    e.Description = description
}

func (e *EC2InstanceLaunchSuccessful) SetEndTime(endTime time.Time) {
    e.EndTime = endTime
}

func (e *EC2InstanceLaunchSuccessful) SetRequestId(requestId string) {
    e.RequestId = requestId
}

func (e *EC2InstanceLaunchSuccessful) SetActivityId(activityId string) {
    e.ActivityId = activityId
}

func (e *EC2InstanceLaunchSuccessful) SetCause(cause string) {
    e.Cause = cause
}

func (e *EC2InstanceLaunchSuccessful) SetAutoScalingGroupName(autoScalingGroupName string) {
    e.AutoScalingGroupName = autoScalingGroupName
}

func (e *EC2InstanceLaunchSuccessful) SetStartTime(startTime time.Time) {
    e.StartTime = startTime
}

func (e *EC2InstanceLaunchSuccessful) SetEC2InstanceId(eC2InstanceId string) {
    e.EC2InstanceId = eC2InstanceId
}

func (e *EC2InstanceLaunchSuccessful) SetStatusCode(statusCode string) {
    e.StatusCode = statusCode
}

func (e *EC2InstanceLaunchSuccessful) SetStatusMessage(statusMessage string) {
    e.StatusMessage = statusMessage
}

func (e *EC2InstanceLaunchSuccessful) SetOrigin(origin string) {
    e.Origin = origin
}

func (e *EC2InstanceLaunchSuccessful) SetDestination(destination string) {
    e.Destination = destination
}
