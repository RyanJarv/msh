package batchjobstatechange

type RetryStrategy struct {
    Attempts float64 `json:"attempts"`
}

func (r *RetryStrategy) SetAttempts(attempts float64) {
    r.Attempts = attempts
}
