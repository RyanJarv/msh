package sagemakermodelbuildingpipelineexecutionstepstatuschange

type Metadata struct {
    ProcessingJob ProcessingJob `json:"processingJob"`
}

func (m *Metadata) SetProcessingJob(processingJob ProcessingJob) {
    m.ProcessingJob = processingJob
}
