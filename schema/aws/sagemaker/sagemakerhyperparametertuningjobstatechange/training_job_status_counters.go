package sagemakerhyperparametertuningjobstatechange

type TrainingJobStatusCounters struct {
    RetryableError float64 `json:"RetryableError"`
    Stopped float64 `json:"Stopped"`
    Completed float64 `json:"Completed"`
    InProgress float64 `json:"InProgress"`
    NonRetryableError float64 `json:"NonRetryableError"`
}

func (t *TrainingJobStatusCounters) SetRetryableError(retryableError float64) {
    t.RetryableError = retryableError
}

func (t *TrainingJobStatusCounters) SetStopped(stopped float64) {
    t.Stopped = stopped
}

func (t *TrainingJobStatusCounters) SetCompleted(completed float64) {
    t.Completed = completed
}

func (t *TrainingJobStatusCounters) SetInProgress(inProgress float64) {
    t.InProgress = inProgress
}

func (t *TrainingJobStatusCounters) SetNonRetryableError(nonRetryableError float64) {
    t.NonRetryableError = nonRetryableError
}
