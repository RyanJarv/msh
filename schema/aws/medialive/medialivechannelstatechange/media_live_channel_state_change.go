package medialivechannelstatechange

type MediaLiveChannelStateChange struct {
    Pipeline string `json:"pipeline,omitempty"`
    Message string `json:"message,omitempty"`
    ChannelArn string `json:"channel_arn,omitempty"`
    PipelinesRunningCount float64 `json:"pipelines_running_count,omitempty"`
    State string `json:"state,omitempty"`
}

func (m *MediaLiveChannelStateChange) SetPipeline(pipeline string) {
    m.Pipeline = pipeline
}

func (m *MediaLiveChannelStateChange) SetMessage(message string) {
    m.Message = message
}

func (m *MediaLiveChannelStateChange) SetChannelArn(channelArn string) {
    m.ChannelArn = channelArn
}

func (m *MediaLiveChannelStateChange) SetPipelinesRunningCount(pipelinesRunningCount float64) {
    m.PipelinesRunningCount = pipelinesRunningCount
}

func (m *MediaLiveChannelStateChange) SetState(state string) {
    m.State = state
}
