package awsapicallviacloudtrail

type RequestParameters struct {
    Input string `json:"input,omitempty"`
    ExecutionArn string `json:"executionArn,omitempty"`
    RoleArn string `json:"roleArn,omitempty"`
    Name string `json:"name,omitempty"`
    Definition string `json:"definition,omitempty"`
    StateMachineArn string `json:"stateMachineArn,omitempty"`
    ActivityArn string `json:"activityArn,omitempty"`
    Tags []interface{} `json:"tags,omitempty"`
}

func (r *RequestParameters) SetInput(input string) {
    r.Input = input
}

func (r *RequestParameters) SetExecutionArn(executionArn string) {
    r.ExecutionArn = executionArn
}

func (r *RequestParameters) SetRoleArn(roleArn string) {
    r.RoleArn = roleArn
}

func (r *RequestParameters) SetName(name string) {
    r.Name = name
}

func (r *RequestParameters) SetDefinition(definition string) {
    r.Definition = definition
}

func (r *RequestParameters) SetStateMachineArn(stateMachineArn string) {
    r.StateMachineArn = stateMachineArn
}

func (r *RequestParameters) SetActivityArn(activityArn string) {
    r.ActivityArn = activityArn
}

func (r *RequestParameters) SetTags(tags []interface{}) {
    r.Tags = tags
}
