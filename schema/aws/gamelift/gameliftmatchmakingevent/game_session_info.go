package gameliftmatchmakingevent

type GameSessionInfo struct {
    Players []GameSessionInfoItem `json:"players"`
    GameSessionArn string `json:"gameSessionArn,omitempty"`
    IpAddress string `json:"ipAddress,omitempty"`
    Port string `json:"port,omitempty"`
}

func (g *GameSessionInfo) SetPlayers(players []GameSessionInfoItem) {
    g.Players = players
}

func (g *GameSessionInfo) SetGameSessionArn(gameSessionArn string) {
    g.GameSessionArn = gameSessionArn
}

func (g *GameSessionInfo) SetIpAddress(ipAddress string) {
    g.IpAddress = ipAddress
}

func (g *GameSessionInfo) SetPort(port string) {
    g.Port = port
}
