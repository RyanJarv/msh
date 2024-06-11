package awsapicallviacloudtrail

type ResponseElements struct {
    PolicyARN string `json:"policyARN,omitempty"`
    FailedScheduledActions []interface{} `json:"failedScheduledActions,omitempty"`
    FailedScheduledUpdateGroupActions []interface{} `json:"failedScheduledUpdateGroupActions,omitempty"`
    Alarms []interface{} `json:"alarms,omitempty"`
}

func (r *ResponseElements) SetPolicyARN(policyARN string) {
    r.PolicyARN = policyARN
}

func (r *ResponseElements) SetFailedScheduledActions(failedScheduledActions []interface{}) {
    r.FailedScheduledActions = failedScheduledActions
}

func (r *ResponseElements) SetFailedScheduledUpdateGroupActions(failedScheduledUpdateGroupActions []interface{}) {
    r.FailedScheduledUpdateGroupActions = failedScheduledUpdateGroupActions
}

func (r *ResponseElements) SetAlarms(alarms []interface{}) {
    r.Alarms = alarms
}
