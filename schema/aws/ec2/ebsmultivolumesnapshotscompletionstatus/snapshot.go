package ebsmultivolumesnapshotscompletionstatus

type Snapshot struct {
    SnapshotId string `json:"snapshot_id,omitempty"`
    Source string `json:"source,omitempty"`
    Status string `json:"status,omitempty"`
}

func (s *Snapshot) SetSnapshotId(snapshotId string) {
    s.SnapshotId = snapshotId
}

func (s *Snapshot) SetSource(source string) {
    s.Source = source
}

func (s *Snapshot) SetStatus(status string) {
    s.Status = status
}
