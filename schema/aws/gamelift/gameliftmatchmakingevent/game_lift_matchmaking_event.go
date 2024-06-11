package gameliftmatchmakingevent

type GameLiftMatchmakingEvent struct {
    GameSessionInfo GameSessionInfo `json:"gameSessionInfo"`
    Message string `json:"message,omitempty"`
    Reason string `json:"reason,omitempty"`
    RuleEvaluationMetrics []RuleEvaluationMetrics `json:"ruleEvaluationMetrics,omitempty"`
    Tickets []Ticket `json:"tickets"`
    GameLiftMatchmakingEventType string `json:"type"`
    EstimatedWaitMillis string `json:"estimatedWaitMillis,omitempty"`
    AcceptanceTimeout float64 `json:"acceptanceTimeout,omitempty"`
    AcceptanceRequired bool `json:"acceptanceRequired,omitempty"`
    MatchId string `json:"matchId,omitempty"`
    CustomEventData string `json:"customEventData,omitempty"`
}

func (g *GameLiftMatchmakingEvent) SetGameSessionInfo(gameSessionInfo GameSessionInfo) {
    g.GameSessionInfo = gameSessionInfo
}

func (g *GameLiftMatchmakingEvent) SetMessage(message string) {
    g.Message = message
}

func (g *GameLiftMatchmakingEvent) SetReason(reason string) {
    g.Reason = reason
}

func (g *GameLiftMatchmakingEvent) SetRuleEvaluationMetrics(ruleEvaluationMetrics []RuleEvaluationMetrics) {
    g.RuleEvaluationMetrics = ruleEvaluationMetrics
}

func (g *GameLiftMatchmakingEvent) SetTickets(tickets []Ticket) {
    g.Tickets = tickets
}

func (g *GameLiftMatchmakingEvent) SetGameLiftMatchmakingEventType(gameLiftMatchmakingEventType string) {
    g.GameLiftMatchmakingEventType = gameLiftMatchmakingEventType
}

func (g *GameLiftMatchmakingEvent) SetEstimatedWaitMillis(estimatedWaitMillis string) {
    g.EstimatedWaitMillis = estimatedWaitMillis
}

func (g *GameLiftMatchmakingEvent) SetAcceptanceTimeout(acceptanceTimeout float64) {
    g.AcceptanceTimeout = acceptanceTimeout
}

func (g *GameLiftMatchmakingEvent) SetAcceptanceRequired(acceptanceRequired bool) {
    g.AcceptanceRequired = acceptanceRequired
}

func (g *GameLiftMatchmakingEvent) SetMatchId(matchId string) {
    g.MatchId = matchId
}

func (g *GameLiftMatchmakingEvent) SetCustomEventData(customEventData string) {
    g.CustomEventData = customEventData
}
