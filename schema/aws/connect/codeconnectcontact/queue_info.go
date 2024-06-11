package codeconnectcontact

type QueueInfo struct {
    QueueArn string `json:"queueArn"`
    QueueType string `json:"queueType"`
}

func (q *QueueInfo) SetQueueArn(queueArn string) {
    q.QueueArn = queueArn
}

func (q *QueueInfo) SetQueueType(queueType string) {
    q.QueueType = queueType
}
