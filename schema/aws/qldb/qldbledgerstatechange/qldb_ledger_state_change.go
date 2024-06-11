package qldbledgerstatechange

type QLDBLedgerStateChange struct {
    LedgerName string `json:"ledgerName"`
    State string `json:"state"`
}

func (q *QLDBLedgerStateChange) SetLedgerName(ledgerName string) {
    q.LedgerName = ledgerName
}

func (q *QLDBLedgerStateChange) SetState(state string) {
    q.State = state
}
