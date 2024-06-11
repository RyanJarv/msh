package roomstatechange

type RoomStateChange struct {
    CalendarEvent Calendar_event `json:"calendar-event"`
    EventType string `json:"event-type"`
    Mode string `json:"mode"`
    Version string `json:"version"`
}

func (r *RoomStateChange) SetCalendarEvent(calendarEvent Calendar_event) {
    r.CalendarEvent = calendarEvent
}

func (r *RoomStateChange) SetEventType(eventType string) {
    r.EventType = eventType
}

func (r *RoomStateChange) SetMode(mode string) {
    r.Mode = mode
}

func (r *RoomStateChange) SetVersion(version string) {
    r.Version = version
}
