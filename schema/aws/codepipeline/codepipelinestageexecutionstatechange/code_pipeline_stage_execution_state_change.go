package codepipelinestageexecutionstatechange

type CodePipelineStageExecutionStateChange struct {
    Pipeline string `json:"pipeline"`
    ExecutionId string `json:"execution-id"`
    Stage string `json:"stage"`
    State string `json:"state"`
    Version string `json:"version"`
}

func (c *CodePipelineStageExecutionStateChange) SetPipeline(pipeline string) {
    c.Pipeline = pipeline
}

func (c *CodePipelineStageExecutionStateChange) SetExecutionId(executionId string) {
    c.ExecutionId = executionId
}

func (c *CodePipelineStageExecutionStateChange) SetStage(stage string) {
    c.Stage = stage
}

func (c *CodePipelineStageExecutionStateChange) SetState(state string) {
    c.State = state
}

func (c *CodePipelineStageExecutionStateChange) SetVersion(version string) {
    c.Version = version
}
