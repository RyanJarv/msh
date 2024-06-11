package datasynclocationstatechange

type DataSyncLocationStateChange struct {
    State []string `json:"State"`
}

func (d *DataSyncLocationStateChange) SetState(state []string) {
    d.State = state
}
