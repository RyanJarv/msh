package awsapicallviacloudtrail

type RequestParameters struct {
    JobName string `json:"jobName"`
    AllocatedCapacity float64 `json:"allocatedCapacity"`
}

func (r *RequestParameters) SetJobName(jobName string) {
    r.JobName = jobName
}

func (r *RequestParameters) SetAllocatedCapacity(allocatedCapacity float64) {
    r.AllocatedCapacity = allocatedCapacity
}
