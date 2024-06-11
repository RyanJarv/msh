package datasynctaskstatechange

type DataSyncTaskStateChange struct {
    State []string `json:"State"`
}

func (d *DataSyncTaskStateChange) SetState(state []string) {
    d.State = state
}
