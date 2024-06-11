package locationgeofenceevent

type LocationGeofenceEvent_Accuracy struct {
    Horizontal float64 `json:"Horizontal,omitempty"`
}

func (l *LocationGeofenceEvent_Accuracy) SetHorizontal(horizontal float64) {
    l.Horizontal = horizontal
}
