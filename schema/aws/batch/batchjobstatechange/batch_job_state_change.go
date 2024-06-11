package batchjobstatechange

type BatchJobStateChange struct {
    Container Container `json:"container"`
    RetryStrategy RetryStrategy `json:"retryStrategy"`
    Parameters interface{} `json:"parameters"`
    JobName string `json:"jobName"`
    JobQueue string `json:"jobQueue"`
    DependsOn []JobDependency `json:"dependsOn"`
    StartedAt int64 `json:"startedAt,omitempty"`
    JobId string `json:"jobId"`
    CreatedAt int64 `json:"createdAt"`
    StoppedAt int64 `json:"stoppedAt,omitempty"`
    StatusReason string `json:"statusReason,omitempty"`
    JobDefinition string `json:"jobDefinition"`
    Status string `json:"status"`
    Attempts []BatchJobStateChangeItem `json:"attempts"`
}

func (b *BatchJobStateChange) SetContainer(container Container) {
    b.Container = container
}

func (b *BatchJobStateChange) SetRetryStrategy(retryStrategy RetryStrategy) {
    b.RetryStrategy = retryStrategy
}

func (b *BatchJobStateChange) SetParameters(parameters interface{}) {
    b.Parameters = parameters
}

func (b *BatchJobStateChange) SetJobName(jobName string) {
    b.JobName = jobName
}

func (b *BatchJobStateChange) SetJobQueue(jobQueue string) {
    b.JobQueue = jobQueue
}

func (b *BatchJobStateChange) SetDependsOn(dependsOn []JobDependency) {
    b.DependsOn = dependsOn
}

func (b *BatchJobStateChange) SetStartedAt(startedAt int64) {
    b.StartedAt = startedAt
}

func (b *BatchJobStateChange) SetJobId(jobId string) {
    b.JobId = jobId
}

func (b *BatchJobStateChange) SetCreatedAt(createdAt int64) {
    b.CreatedAt = createdAt
}

func (b *BatchJobStateChange) SetStoppedAt(stoppedAt int64) {
    b.StoppedAt = stoppedAt
}

func (b *BatchJobStateChange) SetStatusReason(statusReason string) {
    b.StatusReason = statusReason
}

func (b *BatchJobStateChange) SetJobDefinition(jobDefinition string) {
    b.JobDefinition = jobDefinition
}

func (b *BatchJobStateChange) SetStatus(status string) {
    b.Status = status
}

func (b *BatchJobStateChange) SetAttempts(attempts []BatchJobStateChangeItem) {
    b.Attempts = attempts
}
