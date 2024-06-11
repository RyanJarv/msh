package batchjobstatechange

type BatchJobStateChangeItem struct {
    Container Container_1 `json:"container"`
    StoppedAt int64 `json:"stoppedAt"`
    StatusReason string `json:"statusReason"`
    StartedAt int64 `json:"startedAt"`
}

func (b *BatchJobStateChangeItem) SetContainer(container Container_1) {
    b.Container = container
}

func (b *BatchJobStateChangeItem) SetStoppedAt(stoppedAt int64) {
    b.StoppedAt = stoppedAt
}

func (b *BatchJobStateChangeItem) SetStatusReason(statusReason string) {
    b.StatusReason = statusReason
}

func (b *BatchJobStateChangeItem) SetStartedAt(startedAt int64) {
    b.StartedAt = startedAt
}
