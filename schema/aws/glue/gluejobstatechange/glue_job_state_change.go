package gluejobstatechange

type GlueJobStateChange struct {
    JobName string `json:"jobName"`
    Severity string `json:"severity"`
    State string `json:"state"`
    JobRunId string `json:"jobRunId"`
    Message string `json:"message"`
}

func (g *GlueJobStateChange) SetJobName(jobName string) {
    g.JobName = jobName
}

func (g *GlueJobStateChange) SetSeverity(severity string) {
    g.Severity = severity
}

func (g *GlueJobStateChange) SetState(state string) {
    g.State = state
}

func (g *GlueJobStateChange) SetJobRunId(jobRunId string) {
    g.JobRunId = jobRunId
}

func (g *GlueJobStateChange) SetMessage(message string) {
    g.Message = message
}
