package locationdevicepositionevent

import (
    "time"
)


type LocationDevicePositionEvent struct {
    EventType string `json:"EventType,omitempty"`
    TrackerName string `json:"TrackerName,omitempty"`
    DeviceId string `json:"DeviceId,omitempty"`
    SampleTime time.Time `json:"SampleTime,omitempty"`
    ReceivedTime time.Time `json:"ReceivedTime,omitempty"`
    Position []float64 `json:"Position,omitempty"`
    Accuracy LocationDevicePositionEvent_Accuracy `json:"Accuracy,omitempty"`
    PositionProperties map[string]string `json:"PositionProperties,omitempty"`
}

func (l *LocationDevicePositionEvent) SetEventType(eventType string) {
    l.EventType = eventType
}

func (l *LocationDevicePositionEvent) SetTrackerName(trackerName string) {
    l.TrackerName = trackerName
}

func (l *LocationDevicePositionEvent) SetDeviceId(deviceId string) {
    l.DeviceId = deviceId
}

func (l *LocationDevicePositionEvent) SetSampleTime(sampleTime time.Time) {
    l.SampleTime = sampleTime
}

func (l *LocationDevicePositionEvent) SetReceivedTime(receivedTime time.Time) {
    l.ReceivedTime = receivedTime
}

func (l *LocationDevicePositionEvent) SetPosition(position []float64) {
    l.Position = position
}

func (l *LocationDevicePositionEvent) SetAccuracy(accuracy LocationDevicePositionEvent_Accuracy) {
    l.Accuracy = accuracy
}

func (l *LocationDevicePositionEvent) SetPositionProperties(positionProperties map[string]string) {
    l.PositionProperties = positionProperties
}
