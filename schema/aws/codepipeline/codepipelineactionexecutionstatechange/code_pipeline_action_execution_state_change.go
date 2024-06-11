package codepipelineactionexecutionstatechange

type CodePipelineActionExecutionStateChange struct {
    CodePipelineActionExecutionStateChangeType Type `json:"type"`
    Pipeline string `json:"pipeline"`
    ExecutionId string `json:"execution-id"`
    Stage string `json:"stage"`
    Action string `json:"action"`
    State string `json:"state"`
    Region string `json:"region"`
    Version float64 `json:"version"`
    ExecutionResult ExecutionResult `json:"execution-result,omitempty"`
    InputArtifacts InputArtifacts `json:"input-artifacts"`
    OutputArtifacts OutputArtifacts `json:"output-artifacts"`
}

func (c *CodePipelineActionExecutionStateChange) SetCodePipelineActionExecutionStateChangeType(codePipelineActionExecutionStateChangeType Type) {
    c.CodePipelineActionExecutionStateChangeType = codePipelineActionExecutionStateChangeType
}

func (c *CodePipelineActionExecutionStateChange) SetPipeline(pipeline string) {
    c.Pipeline = pipeline
}

func (c *CodePipelineActionExecutionStateChange) SetExecutionId(executionId string) {
    c.ExecutionId = executionId
}

func (c *CodePipelineActionExecutionStateChange) SetStage(stage string) {
    c.Stage = stage
}

func (c *CodePipelineActionExecutionStateChange) SetAction(action string) {
    c.Action = action
}

func (c *CodePipelineActionExecutionStateChange) SetState(state string) {
    c.State = state
}

func (c *CodePipelineActionExecutionStateChange) SetRegion(region string) {
    c.Region = region
}

func (c *CodePipelineActionExecutionStateChange) SetVersion(version float64) {
    c.Version = version
}

func (c *CodePipelineActionExecutionStateChange) SetExecutionResult(executionResult ExecutionResult) {
    c.ExecutionResult = executionResult
}

func (c *CodePipelineActionExecutionStateChange) SetInputArtifacts(inputArtifacts InputArtifacts) {
    c.InputArtifacts = inputArtifacts
}

func (c *CodePipelineActionExecutionStateChange) SetOutputArtifacts(outputArtifacts OutputArtifacts) {
    c.OutputArtifacts = outputArtifacts
}
