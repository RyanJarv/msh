package roomstatechange

import (
    "time"
)


type Calendar_event struct {
    CalendarId string `json:"calendarId"`
    CheckInTime time.Time `json:"checkInTime,omitempty"`
    EndTime time.Time `json:"endTime"`
    Id string `json:"id"`
    ProfileArn string `json:"profileArn"`
    Provider string `json:"provider"`
    ReleaseTime time.Time `json:"releaseTime,omitempty"`
    RoomArn string `json:"roomArn"`
    StartTime time.Time `json:"startTime"`
}

func (c *Calendar_event) SetCalendarId(calendarId string) {
    c.CalendarId = calendarId
}

func (c *Calendar_event) SetCheckInTime(checkInTime time.Time) {
    c.CheckInTime = checkInTime
}

func (c *Calendar_event) SetEndTime(endTime time.Time) {
    c.EndTime = endTime
}

func (c *Calendar_event) SetId(id string) {
    c.Id = id
}

func (c *Calendar_event) SetProfileArn(profileArn string) {
    c.ProfileArn = profileArn
}

func (c *Calendar_event) SetProvider(provider string) {
    c.Provider = provider
}

func (c *Calendar_event) SetReleaseTime(releaseTime time.Time) {
    c.ReleaseTime = releaseTime
}

func (c *Calendar_event) SetRoomArn(roomArn string) {
    c.RoomArn = roomArn
}

func (c *Calendar_event) SetStartTime(startTime time.Time) {
    c.StartTime = startTime
}
