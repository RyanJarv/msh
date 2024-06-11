package simpleworkflowexecutionstatechange

type ExecutionInfo struct {
    Execution Execution `json:"execution"`
    WorkflowType WorkflowType `json:"workflowType"`
    CancelRequested bool `json:"cancelRequested"`
    CloseStatus interface{} `json:"closeStatus"`
    CloseTimestamp interface{} `json:"closeTimestamp"`
    ExecutionStatus string `json:"executionStatus"`
    Parent interface{} `json:"parent"`
    ParentExecutionArn interface{} `json:"parentExecutionArn"`
    StartTimestamp int64 `json:"startTimestamp"`
    TagList interface{} `json:"tagList"`
}

func (e *ExecutionInfo) SetExecution(execution Execution) {
    e.Execution = execution
}

func (e *ExecutionInfo) SetWorkflowType(workflowType WorkflowType) {
    e.WorkflowType = workflowType
}

func (e *ExecutionInfo) SetCancelRequested(cancelRequested bool) {
    e.CancelRequested = cancelRequested
}

func (e *ExecutionInfo) SetCloseStatus(closeStatus interface{}) {
    e.CloseStatus = closeStatus
}

func (e *ExecutionInfo) SetCloseTimestamp(closeTimestamp interface{}) {
    e.CloseTimestamp = closeTimestamp
}

func (e *ExecutionInfo) SetExecutionStatus(executionStatus string) {
    e.ExecutionStatus = executionStatus
}

func (e *ExecutionInfo) SetParent(parent interface{}) {
    e.Parent = parent
}

func (e *ExecutionInfo) SetParentExecutionArn(parentExecutionArn interface{}) {
    e.ParentExecutionArn = parentExecutionArn
}

func (e *ExecutionInfo) SetStartTimestamp(startTimestamp int64) {
    e.StartTimestamp = startTimestamp
}

func (e *ExecutionInfo) SetTagList(tagList interface{}) {
    e.TagList = tagList
}
