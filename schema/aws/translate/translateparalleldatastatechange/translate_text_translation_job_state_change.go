package translateparalleldatastatechange

type TranslateTextTranslationJobStateChange struct {
    JobId string `json:"jobId"`
    JobStatus string `json:"jobStatus"`
}

func (t *TranslateTextTranslationJobStateChange) SetJobId(jobId string) {
    t.JobId = jobId
}

func (t *TranslateTextTranslationJobStateChange) SetJobStatus(jobStatus string) {
    t.JobStatus = jobStatus
}
