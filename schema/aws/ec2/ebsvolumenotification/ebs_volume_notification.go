package ebsvolumenotification

type EBSVolumeNotification struct {
    Result string `json:"result"`
    Cause string `json:"cause"`
    Event string `json:"event"`
    RequestId string `json:"request-id"`
}

func (e *EBSVolumeNotification) SetResult(result string) {
    e.Result = result
}

func (e *EBSVolumeNotification) SetCause(cause string) {
    e.Cause = cause
}

func (e *EBSVolumeNotification) SetEvent(event string) {
    e.Event = event
}

func (e *EBSVolumeNotification) SetRequestId(requestId string) {
    e.RequestId = requestId
}
