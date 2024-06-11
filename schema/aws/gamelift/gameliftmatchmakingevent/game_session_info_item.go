package gameliftmatchmakingevent

type GameSessionInfoItem struct {
    PlayerId string `json:"playerId"`
    Team string `json:"team,omitempty"`
}

func (g *GameSessionInfoItem) SetPlayerId(playerId string) {
    g.PlayerId = playerId
}

func (g *GameSessionInfoItem) SetTeam(team string) {
    g.Team = team
}
