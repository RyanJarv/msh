package voiceidbatchfraudsterregistrationaction

type RegistrationConfig struct {
    DuplicateRegistrationAction string `json:"duplicateRegistrationAction,omitempty"`
    FraudsterSimilarityThreshold float64 `json:"fraudsterSimilarityThreshold,omitempty"`
    WatchlistIds []string `json:"watchlistIds,omitempty"`
}

func (r *RegistrationConfig) SetDuplicateRegistrationAction(duplicateRegistrationAction string) {
    r.DuplicateRegistrationAction = duplicateRegistrationAction
}

func (r *RegistrationConfig) SetFraudsterSimilarityThreshold(fraudsterSimilarityThreshold float64) {
    r.FraudsterSimilarityThreshold = fraudsterSimilarityThreshold
}

func (r *RegistrationConfig) SetWatchlistIds(watchlistIds []string) {
    r.WatchlistIds = watchlistIds
}
