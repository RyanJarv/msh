package voiceidupdatesessionaction

type FraudDetectionConfiguration struct {
    RiskThreshold float64 `json:"riskThreshold,omitempty"`
    WatchlistId string `json:"watchlistId,omitempty"`
}

func (f *FraudDetectionConfiguration) SetRiskThreshold(riskThreshold float64) {
    f.RiskThreshold = riskThreshold
}

func (f *FraudDetectionConfiguration) SetWatchlistId(watchlistId string) {
    f.WatchlistId = watchlistId
}
