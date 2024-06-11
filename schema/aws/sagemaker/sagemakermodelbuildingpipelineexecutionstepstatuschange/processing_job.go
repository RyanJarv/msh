package sagemakermodelbuildingpipelineexecutionstepstatuschange

type ProcessingJob struct {
    Arn string `json:"arn"`
}

func (p *ProcessingJob) SetArn(arn string) {
    p.Arn = arn
}
