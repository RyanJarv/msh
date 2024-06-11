package healthlakeimportjobstatuschange

type HealthLakeImportJobStatusChange struct {
    InputDataConfig InputDataConfig `json:"inputDataConfig,omitempty"`
    DatastoreId string `json:"datastoreId,omitempty"`
    JobId string `json:"jobId,omitempty"`
    JobStatus string `json:"jobStatus,omitempty"`
    SubmitTime string `json:"submitTime,omitempty"`
}

func (h *HealthLakeImportJobStatusChange) SetInputDataConfig(inputDataConfig InputDataConfig) {
    h.InputDataConfig = inputDataConfig
}

func (h *HealthLakeImportJobStatusChange) SetDatastoreId(datastoreId string) {
    h.DatastoreId = datastoreId
}

func (h *HealthLakeImportJobStatusChange) SetJobId(jobId string) {
    h.JobId = jobId
}

func (h *HealthLakeImportJobStatusChange) SetJobStatus(jobStatus string) {
    h.JobStatus = jobStatus
}

func (h *HealthLakeImportJobStatusChange) SetSubmitTime(submitTime string) {
    h.SubmitTime = submitTime
}
