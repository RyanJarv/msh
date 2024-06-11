package medialivemultiplexstatechange

type MediaLiveMultiplexStateChange struct {
    Message string `json:"message"`
    Pipeline string `json:"pipeline"`
    State string `json:"state"`
}

func (m *MediaLiveMultiplexStateChange) SetMessage(message string) {
    m.Message = message
}

func (m *MediaLiveMultiplexStateChange) SetPipeline(pipeline string) {
    m.Pipeline = pipeline
}

func (m *MediaLiveMultiplexStateChange) SetState(state string) {
    m.State = state
}
