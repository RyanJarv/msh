package ebssnapshotnotification

import (
    "time"
)


type EBSSnapshotNotification struct {
    Result string `json:"result"`
    EndTime time.Time `json:"EndTime"`
    SnapshotId string `json:"snapshot_id"`
    StartTime time.Time `json:"StartTime"`
    Cause string `json:"cause"`
    Source string `json:"source"`
    Event string `json:"event"`
    RequestId string `json:"request-id"`
}

func (e *EBSSnapshotNotification) SetResult(result string) {
    e.Result = result
}

func (e *EBSSnapshotNotification) SetEndTime(endTime time.Time) {
    e.EndTime = endTime
}

func (e *EBSSnapshotNotification) SetSnapshotId(snapshotId string) {
    e.SnapshotId = snapshotId
}

func (e *EBSSnapshotNotification) SetStartTime(startTime time.Time) {
    e.StartTime = startTime
}

func (e *EBSSnapshotNotification) SetCause(cause string) {
    e.Cause = cause
}

func (e *EBSSnapshotNotification) SetSource(source string) {
    e.Source = source
}

func (e *EBSSnapshotNotification) SetEvent(event string) {
    e.Event = event
}

func (e *EBSSnapshotNotification) SetRequestId(requestId string) {
    e.RequestId = requestId
}
