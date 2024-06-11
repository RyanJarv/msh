package awsapicallviacloudtrail

type RequestParameters struct {
    PolicyDetails PolicyDetails `json:"PolicyDetails,omitempty"`
    Description string `json:"Description,omitempty"`
    ExecutionRoleArn string `json:"ExecutionRoleArn,omitempty"`
    State string `json:"State,omitempty"`
    PolicyId string `json:"policyId,omitempty"`
}

func (r *RequestParameters) SetPolicyDetails(policyDetails PolicyDetails) {
    r.PolicyDetails = policyDetails
}

func (r *RequestParameters) SetDescription(description string) {
    r.Description = description
}

func (r *RequestParameters) SetExecutionRoleArn(executionRoleArn string) {
    r.ExecutionRoleArn = executionRoleArn
}

func (r *RequestParameters) SetState(state string) {
    r.State = state
}

func (r *RequestParameters) SetPolicyId(policyId string) {
    r.PolicyId = policyId
}
