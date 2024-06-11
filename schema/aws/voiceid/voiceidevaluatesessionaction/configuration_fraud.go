package voiceidevaluatesessionaction

type Configuration_Fraud struct {
    RiskThreshold float64 `json:"riskThreshold,omitempty"`
    WatchlistId string `json:"watchlistId,omitempty"`
}

func (c *Configuration_Fraud) SetRiskThreshold(riskThreshold float64) {
    c.RiskThreshold = riskThreshold
}

func (c *Configuration_Fraud) SetWatchlistId(watchlistId string) {
    c.WatchlistId = watchlistId
}
