package simpleworkflowexecutionstatechange

type Execution struct {
    RunId string `json:"runId"`
    WorkflowId string `json:"workflowId"`
}

func (e *Execution) SetRunId(runId string) {
    e.RunId = runId
}

func (e *Execution) SetWorkflowId(workflowId string) {
    e.WorkflowId = workflowId
}
