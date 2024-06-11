package voiceidbatchspeakerenrollmentaction

type FraudDetectionConfig struct {
    FraudDetectionAction string `json:"fraudDetectionAction,omitempty"`
    FraudDetectionThreshold float64 `json:"fraudDetectionThreshold,omitempty"`
    WatchlistIds []string `json:"watchlistIds,omitempty"`
}

func (f *FraudDetectionConfig) SetFraudDetectionAction(fraudDetectionAction string) {
    f.FraudDetectionAction = fraudDetectionAction
}

func (f *FraudDetectionConfig) SetFraudDetectionThreshold(fraudDetectionThreshold float64) {
    f.FraudDetectionThreshold = fraudDetectionThreshold
}

func (f *FraudDetectionConfig) SetWatchlistIds(watchlistIds []string) {
    f.WatchlistIds = watchlistIds
}
