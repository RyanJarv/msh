package streamhealthchange

type IVSStreamHealthChange struct {
    ChannelName string `json:"channel_name"`
    EventName string `json:"event_name"`
    StreamId string `json:"stream_id"`
}

func (i *IVSStreamHealthChange) SetChannelName(channelName string) {
    i.ChannelName = channelName
}

func (i *IVSStreamHealthChange) SetEventName(eventName string) {
    i.EventName = eventName
}

func (i *IVSStreamHealthChange) SetStreamId(streamId string) {
    i.StreamId = streamId
}
