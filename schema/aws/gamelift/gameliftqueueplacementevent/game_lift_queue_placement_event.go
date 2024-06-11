package gameliftqueueplacementevent

import (
    "time"
)


type GameLiftQueuePlacementEvent struct {
    DnsName string `json:"dnsName"`
    EndTime time.Time `json:"endTime"`
    GameSessionArn string `json:"gameSessionArn"`
    GameSessionRegion string `json:"gameSessionRegion"`
    IpAddress string `json:"ipAddress"`
    PlacedPlayerSessions []GameLiftQueuePlacementEventItem `json:"placedPlayerSessions"`
    PlacementId string `json:"placementId"`
    Port string `json:"port"`
    StartTime time.Time `json:"startTime"`
    GameLiftQueuePlacementEventType string `json:"type"`
}

func (g *GameLiftQueuePlacementEvent) SetDnsName(dnsName string) {
    g.DnsName = dnsName
}

func (g *GameLiftQueuePlacementEvent) SetEndTime(endTime time.Time) {
    g.EndTime = endTime
}

func (g *GameLiftQueuePlacementEvent) SetGameSessionArn(gameSessionArn string) {
    g.GameSessionArn = gameSessionArn
}

func (g *GameLiftQueuePlacementEvent) SetGameSessionRegion(gameSessionRegion string) {
    g.GameSessionRegion = gameSessionRegion
}

func (g *GameLiftQueuePlacementEvent) SetIpAddress(ipAddress string) {
    g.IpAddress = ipAddress
}

func (g *GameLiftQueuePlacementEvent) SetPlacedPlayerSessions(placedPlayerSessions []GameLiftQueuePlacementEventItem) {
    g.PlacedPlayerSessions = placedPlayerSessions
}

func (g *GameLiftQueuePlacementEvent) SetPlacementId(placementId string) {
    g.PlacementId = placementId
}

func (g *GameLiftQueuePlacementEvent) SetPort(port string) {
    g.Port = port
}

func (g *GameLiftQueuePlacementEvent) SetStartTime(startTime time.Time) {
    g.StartTime = startTime
}

func (g *GameLiftQueuePlacementEvent) SetGameLiftQueuePlacementEventType(gameLiftQueuePlacementEventType string) {
    g.GameLiftQueuePlacementEventType = gameLiftQueuePlacementEventType
}
