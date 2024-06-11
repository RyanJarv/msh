package codepipelinepipelineexecutionstatechange

type CodePipelinePipelineExecutionStateChange_executiontrigger struct {
    TriggerType string `json:"trigger-type"`
    TriggerDetail string `json:"trigger-detail"`
}

func (c *CodePipelinePipelineExecutionStateChange_executiontrigger) SetTriggerType(triggerType string) {
    c.TriggerType = triggerType
}

func (c *CodePipelinePipelineExecutionStateChange_executiontrigger) SetTriggerDetail(triggerDetail string) {
    c.TriggerDetail = triggerDetail
}
