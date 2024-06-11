package medialivemultiplexalert

type MediaLiveMultiplexAlert struct {
    AlarmState string `json:"alarm_state"`
    AlertType string `json:"alert_type"`
    Message string `json:"message"`
    Pipeline string `json:"pipeline"`
}

func (m *MediaLiveMultiplexAlert) SetAlarmState(alarmState string) {
    m.AlarmState = alarmState
}

func (m *MediaLiveMultiplexAlert) SetAlertType(alertType string) {
    m.AlertType = alertType
}

func (m *MediaLiveMultiplexAlert) SetMessage(message string) {
    m.Message = message
}

func (m *MediaLiveMultiplexAlert) SetPipeline(pipeline string) {
    m.Pipeline = pipeline
}
