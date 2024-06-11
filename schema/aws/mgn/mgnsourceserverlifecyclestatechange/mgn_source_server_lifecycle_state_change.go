package mgnsourceserverlifecyclestatechange

type MGNSourceServerLifecycleStateChange struct {
    State string `json:"state"`
}

func (m *MGNSourceServerLifecycleStateChange) SetState(state string) {
    m.State = state
}
