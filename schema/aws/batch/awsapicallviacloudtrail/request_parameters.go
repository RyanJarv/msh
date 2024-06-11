package awsapicallviacloudtrail

type RequestParameters struct {
    ContainerOverrides ContainerOverrides `json:"containerOverrides,omitempty"`
    JobName string `json:"jobName"`
    JobQueue string `json:"jobQueue"`
    JobDefinition string `json:"jobDefinition"`
}

func (r *RequestParameters) SetContainerOverrides(containerOverrides ContainerOverrides) {
    r.ContainerOverrides = containerOverrides
}

func (r *RequestParameters) SetJobName(jobName string) {
    r.JobName = jobName
}

func (r *RequestParameters) SetJobQueue(jobQueue string) {
    r.JobQueue = jobQueue
}

func (r *RequestParameters) SetJobDefinition(jobDefinition string) {
    r.JobDefinition = jobDefinition
}
