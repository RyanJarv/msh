package codepipelineactionexecutionstatechange

type ExecutionResult struct {
    ExternalExecutionId string `json:"external-execution-id,omitempty"`
    ExternalExecutionUrl string `json:"external-execution-url,omitempty"`
    ExternalExecutionSummary string `json:"external-execution-summary,omitempty"`
}

func (e *ExecutionResult) SetExternalExecutionId(externalExecutionId string) {
    e.ExternalExecutionId = externalExecutionId
}

func (e *ExecutionResult) SetExternalExecutionUrl(externalExecutionUrl string) {
    e.ExternalExecutionUrl = externalExecutionUrl
}

func (e *ExecutionResult) SetExternalExecutionSummary(externalExecutionSummary string) {
    e.ExternalExecutionSummary = externalExecutionSummary
}
