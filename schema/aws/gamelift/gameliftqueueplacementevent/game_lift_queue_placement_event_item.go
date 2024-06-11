package gameliftqueueplacementevent

type GameLiftQueuePlacementEventItem struct {
    PlayerId string `json:"playerId"`
    PlayerSessionId string `json:"playerSessionId"`
}

func (g *GameLiftQueuePlacementEventItem) SetPlayerId(playerId string) {
    g.PlayerId = playerId
}

func (g *GameLiftQueuePlacementEventItem) SetPlayerSessionId(playerSessionId string) {
    g.PlayerSessionId = playerSessionId
}
