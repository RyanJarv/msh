package cloudformationhookinvocationprogress

type CloudFormationHookInvocationProgress struct {
    HookDetail Hook_detail `json:"hook-detail"`
    Result Result `json:"result"`
    TargetDetail Target_detail `json:"target-detail"`
    Status string `json:"status"`
    StatusReason string `json:"status-reason"`
}

func (c *CloudFormationHookInvocationProgress) SetHookDetail(hookDetail Hook_detail) {
    c.HookDetail = hookDetail
}

func (c *CloudFormationHookInvocationProgress) SetResult(result Result) {
    c.Result = result
}

func (c *CloudFormationHookInvocationProgress) SetTargetDetail(targetDetail Target_detail) {
    c.TargetDetail = targetDetail
}

func (c *CloudFormationHookInvocationProgress) SetStatus(status string) {
    c.Status = status
}

func (c *CloudFormationHookInvocationProgress) SetStatusReason(statusReason string) {
    c.StatusReason = statusReason
}
