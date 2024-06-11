package streamstatechange

type IVSStreamStateChange struct {
    ChannelName string `json:"channel_name"`
    EventName string `json:"event_name"`
    StreamId string `json:"stream_id"`
}

func (i *IVSStreamStateChange) SetChannelName(channelName string) {
    i.ChannelName = channelName
}

func (i *IVSStreamStateChange) SetEventName(eventName string) {
    i.EventName = eventName
}

func (i *IVSStreamStateChange) SetStreamId(streamId string) {
    i.StreamId = streamId
}
