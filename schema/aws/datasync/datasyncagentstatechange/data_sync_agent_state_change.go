package datasyncagentstatechange

type DataSyncAgentStateChange struct {
    State []string `json:"State"`
}

func (d *DataSyncAgentStateChange) SetState(state []string) {
    d.State = state
}
