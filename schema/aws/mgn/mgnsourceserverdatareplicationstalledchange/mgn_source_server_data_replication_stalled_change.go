package mgnsourceserverdatareplicationstalledchange

type MGNSourceServerDataReplicationStalledChange struct {
    State string `json:"state"`
}

func (m *MGNSourceServerDataReplicationStalledChange) SetState(state string) {
    m.State = state
}
