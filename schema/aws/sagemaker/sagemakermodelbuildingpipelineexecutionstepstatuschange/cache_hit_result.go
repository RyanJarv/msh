package sagemakermodelbuildingpipelineexecutionstepstatuschange

type CacheHitResult struct {
    SourcePipelineExecutionArn string `json:"sourcePipelineExecutionArn"`
}

func (c *CacheHitResult) SetSourcePipelineExecutionArn(sourcePipelineExecutionArn string) {
    c.SourcePipelineExecutionArn = sourcePipelineExecutionArn
}
