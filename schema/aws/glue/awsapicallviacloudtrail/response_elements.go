package awsapicallviacloudtrail

type ResponseElements struct {
    JobRunId string `json:"jobRunId"`
}

func (r *ResponseElements) SetJobRunId(jobRunId string) {
    r.JobRunId = jobRunId
}
