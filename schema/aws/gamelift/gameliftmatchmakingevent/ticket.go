package gameliftmatchmakingevent

import (
    "time"
)


type Ticket struct {
    Players []GameSessionInfoItem `json:"players"`
    StartTime time.Time `json:"startTime"`
    TicketId string `json:"ticketId"`
}

func (t *Ticket) SetPlayers(players []GameSessionInfoItem) {
    t.Players = players
}

func (t *Ticket) SetStartTime(startTime time.Time) {
    t.StartTime = startTime
}

func (t *Ticket) SetTicketId(ticketId string) {
    t.TicketId = ticketId
}
