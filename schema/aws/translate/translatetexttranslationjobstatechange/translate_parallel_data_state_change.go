package translatetexttranslationjobstatechange

import (
    "time"
)


type TranslateParallelDataStateChange struct {
    LatestUpdateAttemptAt time.Time `json:"latestUpdateAttemptAt"`
    LatestUpdateAttemptStatus string `json:"latestUpdateAttemptStatus"`
    Name string `json:"name"`
    Operation string `json:"operation"`
    Status string `json:"status"`
}

func (t *TranslateParallelDataStateChange) SetLatestUpdateAttemptAt(latestUpdateAttemptAt time.Time) {
    t.LatestUpdateAttemptAt = latestUpdateAttemptAt
}

func (t *TranslateParallelDataStateChange) SetLatestUpdateAttemptStatus(latestUpdateAttemptStatus string) {
    t.LatestUpdateAttemptStatus = latestUpdateAttemptStatus
}

func (t *TranslateParallelDataStateChange) SetName(name string) {
    t.Name = name
}

func (t *TranslateParallelDataStateChange) SetOperation(operation string) {
    t.Operation = operation
}

func (t *TranslateParallelDataStateChange) SetStatus(status string) {
    t.Status = status
}
