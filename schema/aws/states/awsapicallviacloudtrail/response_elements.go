package awsapicallviacloudtrail

type ResponseElements struct {
    UpdateDate string `json:"updateDate,omitempty"`
    ExecutionArn string `json:"executionArn,omitempty"`
    CreationDate string `json:"creationDate,omitempty"`
    StateMachineArn string `json:"stateMachineArn,omitempty"`
    ActivityArn string `json:"activityArn,omitempty"`
    StartDate string `json:"startDate,omitempty"`
    StopDate string `json:"stopDate,omitempty"`
}

func (r *ResponseElements) SetUpdateDate(updateDate string) {
    r.UpdateDate = updateDate
}

func (r *ResponseElements) SetExecutionArn(executionArn string) {
    r.ExecutionArn = executionArn
}

func (r *ResponseElements) SetCreationDate(creationDate string) {
    r.CreationDate = creationDate
}

func (r *ResponseElements) SetStateMachineArn(stateMachineArn string) {
    r.StateMachineArn = stateMachineArn
}

func (r *ResponseElements) SetActivityArn(activityArn string) {
    r.ActivityArn = activityArn
}

func (r *ResponseElements) SetStartDate(startDate string) {
    r.StartDate = startDate
}

func (r *ResponseElements) SetStopDate(stopDate string) {
    r.StopDate = stopDate
}
