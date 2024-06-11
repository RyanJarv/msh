package medialivechannelinputchange

type MediaLiveChannelInputChange struct {
    ActiveInputAttachmentName string `json:"active_input_attachment_name"`
    ActiveInputSwitchActionName string `json:"active_input_switch_action_name"`
    ChannelArn string `json:"channel_arn"`
    Message string `json:"message"`
    Pipeline string `json:"pipeline"`
}

func (m *MediaLiveChannelInputChange) SetActiveInputAttachmentName(activeInputAttachmentName string) {
    m.ActiveInputAttachmentName = activeInputAttachmentName
}

func (m *MediaLiveChannelInputChange) SetActiveInputSwitchActionName(activeInputSwitchActionName string) {
    m.ActiveInputSwitchActionName = activeInputSwitchActionName
}

func (m *MediaLiveChannelInputChange) SetChannelArn(channelArn string) {
    m.ChannelArn = channelArn
}

func (m *MediaLiveChannelInputChange) SetMessage(message string) {
    m.Message = message
}

func (m *MediaLiveChannelInputChange) SetPipeline(pipeline string) {
    m.Pipeline = pipeline
}
