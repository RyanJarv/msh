package ebsmultivolumesnapshotscompletionstatus

type EBSMultiVolumeSnapshotsCompletionStatus struct {
    Cause string `json:"cause,omitempty"`
    EndTime string `json:"endTime,omitempty"`
    Event string `json:"event,omitempty"`
    RequestId string `json:"request-id,omitempty"`
    Result string `json:"result,omitempty"`
    Snapshots []Snapshot `json:"snapshots,omitempty"`
    StartTime string `json:"startTime,omitempty"`
}

func (e *EBSMultiVolumeSnapshotsCompletionStatus) SetCause(cause string) {
    e.Cause = cause
}

func (e *EBSMultiVolumeSnapshotsCompletionStatus) SetEndTime(endTime string) {
    e.EndTime = endTime
}

func (e *EBSMultiVolumeSnapshotsCompletionStatus) SetEvent(event string) {
    e.Event = event
}

func (e *EBSMultiVolumeSnapshotsCompletionStatus) SetRequestId(requestId string) {
    e.RequestId = requestId
}

func (e *EBSMultiVolumeSnapshotsCompletionStatus) SetResult(result string) {
    e.Result = result
}

func (e *EBSMultiVolumeSnapshotsCompletionStatus) SetSnapshots(snapshots []Snapshot) {
    e.Snapshots = snapshots
}

func (e *EBSMultiVolumeSnapshotsCompletionStatus) SetStartTime(startTime string) {
    e.StartTime = startTime
}
