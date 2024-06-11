package gluejobrunstatus

type NotificationCondition struct {
    NotifyDelayAfter float64 `json:"NotifyDelayAfter"`
}

func (n *NotificationCondition) SetNotifyDelayAfter(notifyDelayAfter float64) {
    n.NotifyDelayAfter = notifyDelayAfter
}
