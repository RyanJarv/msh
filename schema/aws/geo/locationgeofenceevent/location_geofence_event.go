package locationgeofenceevent

import (
    "time"
)


type LocationGeofenceEvent struct {
    DeviceId string `json:"DeviceId"`
    EventType string `json:"EventType"`
    GeofenceId string `json:"GeofenceId"`
    Position []float64 `json:"Position"`
    SampleTime time.Time `json:"SampleTime"`
    Accuracy LocationGeofenceEvent_Accuracy `json:"Accuracy,omitempty"`
    PositionProperties map[string]string `json:"PositionProperties,omitempty"`
}

func (l *LocationGeofenceEvent) SetDeviceId(deviceId string) {
    l.DeviceId = deviceId
}

func (l *LocationGeofenceEvent) SetEventType(eventType string) {
    l.EventType = eventType
}

func (l *LocationGeofenceEvent) SetGeofenceId(geofenceId string) {
    l.GeofenceId = geofenceId
}

func (l *LocationGeofenceEvent) SetPosition(position []float64) {
    l.Position = position
}

func (l *LocationGeofenceEvent) SetSampleTime(sampleTime time.Time) {
    l.SampleTime = sampleTime
}

func (l *LocationGeofenceEvent) SetAccuracy(accuracy LocationGeofenceEvent_Accuracy) {
    l.Accuracy = accuracy
}

func (l *LocationGeofenceEvent) SetPositionProperties(positionProperties map[string]string) {
    l.PositionProperties = positionProperties
}
