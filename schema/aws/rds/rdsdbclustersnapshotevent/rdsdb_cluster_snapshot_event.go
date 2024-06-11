package rdsdbclustersnapshotevent

import (
    "time"
)


type RDSDBClusterSnapshotEvent struct {
    SourceArn string `json:"SourceArn,omitempty"`
    Message string `json:"Message,omitempty"`
    SourceType string `json:"SourceType,omitempty"`
    SourceIdentifier string `json:"SourceIdentifier,omitempty"`
    EventCategories []string `json:"EventCategories,omitempty"`
    Date time.Time `json:"Date,omitempty"`
    Tags map[string]string `json:"Tags,omitempty"`
}

func (r *RDSDBClusterSnapshotEvent) SetSourceArn(sourceArn string) {
    r.SourceArn = sourceArn
}

func (r *RDSDBClusterSnapshotEvent) SetMessage(message string) {
    r.Message = message
}

func (r *RDSDBClusterSnapshotEvent) SetSourceType(sourceType string) {
    r.SourceType = sourceType
}

func (r *RDSDBClusterSnapshotEvent) SetSourceIdentifier(sourceIdentifier string) {
    r.SourceIdentifier = sourceIdentifier
}

func (r *RDSDBClusterSnapshotEvent) SetEventCategories(eventCategories []string) {
    r.EventCategories = eventCategories
}

func (r *RDSDBClusterSnapshotEvent) SetDate(date time.Time) {
    r.Date = date
}

func (r *RDSDBClusterSnapshotEvent) SetTags(tags map[string]string) {
    r.Tags = tags
}
