package datasynctaskexecutionstatechange

type DataSyncTaskExecutionStateChange struct {
    State []string `json:"State"`
}

func (d *DataSyncTaskExecutionStateChange) SetState(state []string) {
    d.State = state
}
