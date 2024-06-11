package calendarstatechange

import (
    "time"
)


type CalendarStateChange struct {
    AtTime time.Time `json:"atTime"`
    NextTransitionTime time.Time `json:"nextTransitionTime"`
    State string `json:"state"`
}

func (c *CalendarStateChange) SetAtTime(atTime time.Time) {
    c.AtTime = atTime
}

func (c *CalendarStateChange) SetNextTransitionTime(nextTransitionTime time.Time) {
    c.NextTransitionTime = nextTransitionTime
}

func (c *CalendarStateChange) SetState(state string) {
    c.State = state
}
