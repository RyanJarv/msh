package mgnsourceserverlaunchresult

type MGNSourceServerLaunchResult struct {
    JobId string `json:"job-id"`
    State string `json:"state"`
}

func (m *MGNSourceServerLaunchResult) SetJobId(jobId string) {
    m.JobId = jobId
}

func (m *MGNSourceServerLaunchResult) SetState(state string) {
    m.State = state
}
