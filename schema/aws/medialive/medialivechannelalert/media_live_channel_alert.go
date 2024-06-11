package medialivechannelalert

type MediaLiveChannelAlert struct {
    AlarmId string `json:"alarm_id,omitempty"`
    AlarmState string `json:"alarm_state,omitempty"`
    AlertType string `json:"alert_type,omitempty"`
    ChannelArn string `json:"channel_arn,omitempty"`
    Message string `json:"message,omitempty"`
    Pipeline string `json:"pipeline,omitempty"`
}

func (m *MediaLiveChannelAlert) SetAlarmId(alarmId string) {
    m.AlarmId = alarmId
}

func (m *MediaLiveChannelAlert) SetAlarmState(alarmState string) {
    m.AlarmState = alarmState
}

func (m *MediaLiveChannelAlert) SetAlertType(alertType string) {
    m.AlertType = alertType
}

func (m *MediaLiveChannelAlert) SetChannelArn(channelArn string) {
    m.ChannelArn = channelArn
}

func (m *MediaLiveChannelAlert) SetMessage(message string) {
    m.Message = message
}

func (m *MediaLiveChannelAlert) SetPipeline(pipeline string) {
    m.Pipeline = pipeline
}
