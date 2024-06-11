package locationdevicepositionevent

type LocationDevicePositionEvent_Accuracy struct {
    Horizontal float64 `json:"Horizontal,omitempty"`
}

func (l *LocationDevicePositionEvent_Accuracy) SetHorizontal(horizontal float64) {
    l.Horizontal = horizontal
}
