package awsapicallviacloudtrail

type ResponseElements struct {
    JobName string `json:"jobName"`
    JobId string `json:"jobId"`
}

func (r *ResponseElements) SetJobName(jobName string) {
    r.JobName = jobName
}

func (r *ResponseElements) SetJobId(jobId string) {
    r.JobId = jobId
}
