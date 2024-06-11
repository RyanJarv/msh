package ec2instancelaunchunsuccessful

import (
    "time"
)


type EC2InstanceLaunchUnsuccessful struct {
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

func (e *EC2InstanceLaunchUnsuccessful) SetDetails(details Details) {
    e.Details = details
}

func (e *EC2InstanceLaunchUnsuccessful) SetDescription(description string) {
    e.Description = description
}

func (e *EC2InstanceLaunchUnsuccessful) SetEndTime(endTime time.Time) {
    e.EndTime = endTime
}

func (e *EC2InstanceLaunchUnsuccessful) SetRequestId(requestId string) {
    e.RequestId = requestId
}

func (e *EC2InstanceLaunchUnsuccessful) SetActivityId(activityId string) {
    e.ActivityId = activityId
}

func (e *EC2InstanceLaunchUnsuccessful) SetCause(cause string) {
    e.Cause = cause
}

func (e *EC2InstanceLaunchUnsuccessful) SetAutoScalingGroupName(autoScalingGroupName string) {
    e.AutoScalingGroupName = autoScalingGroupName
}

func (e *EC2InstanceLaunchUnsuccessful) SetStartTime(startTime time.Time) {
    e.StartTime = startTime
}

func (e *EC2InstanceLaunchUnsuccessful) SetEC2InstanceId(eC2InstanceId string) {
    e.EC2InstanceId = eC2InstanceId
}

func (e *EC2InstanceLaunchUnsuccessful) SetStatusCode(statusCode string) {
    e.StatusCode = statusCode
}

func (e *EC2InstanceLaunchUnsuccessful) SetStatusMessage(statusMessage string) {
    e.StatusMessage = statusMessage
}

func (e *EC2InstanceLaunchUnsuccessful) SetOrigin(origin string) {
    e.Origin = origin
}

func (e *EC2InstanceLaunchUnsuccessful) SetDestination(destination string) {
    e.Destination = destination
}
