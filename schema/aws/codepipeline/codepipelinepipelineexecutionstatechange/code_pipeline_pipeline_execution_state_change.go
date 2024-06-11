package codepipelinepipelineexecutionstatechange

type CodePipelinePipelineExecutionStateChange struct {
    Pipeline string `json:"pipeline"`
    ExecutionId string `json:"execution-id"`
    State string `json:"state"`
    Version string `json:"version"`
    ExecutionTrigger CodePipelinePipelineExecutionStateChange_executiontrigger `json:"execution-trigger,omitempty"`
}

func (c *CodePipelinePipelineExecutionStateChange) SetPipeline(pipeline string) {
    c.Pipeline = pipeline
}

func (c *CodePipelinePipelineExecutionStateChange) SetExecutionId(executionId string) {
    c.ExecutionId = executionId
}

func (c *CodePipelinePipelineExecutionStateChange) SetState(state string) {
    c.State = state
}

func (c *CodePipelinePipelineExecutionStateChange) SetVersion(version string) {
    c.Version = version
}

func (c *CodePipelinePipelineExecutionStateChange) SetExecutionTrigger(executionTrigger CodePipelinePipelineExecutionStateChange_executiontrigger) {
    c.ExecutionTrigger = executionTrigger
}
