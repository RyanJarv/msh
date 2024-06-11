package batchjobstatechange

type JobDependency struct {
    JobId string `json:"jobId,omitempty"`
    JobDependencyType string `json:"type,omitempty"`
}

func (j *JobDependency) SetJobId(jobId string) {
    j.JobId = jobId
}

func (j *JobDependency) SetJobDependencyType(jobDependencyType string) {
    j.JobDependencyType = jobDependencyType
}
