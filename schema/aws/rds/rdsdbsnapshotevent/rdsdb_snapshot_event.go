package rdsdbsnapshotevent

import (
    "time"
)


type RDSDBSnapshotEvent struct {
    SourceArn string `json:"SourceArn,omitempty"`
    Message string `json:"Message,omitempty"`
    SourceType string `json:"SourceType,omitempty"`
    SourceIdentifier string `json:"SourceIdentifier,omitempty"`
    EventCategories []string `json:"EventCategories,omitempty"`
    Date time.Time `json:"Date,omitempty"`
    Tags map[string]string `json:"Tags,omitempty"`
}

func (r *RDSDBSnapshotEvent) SetSourceArn(sourceArn string) {
    r.SourceArn = sourceArn
}

func (r *RDSDBSnapshotEvent) SetMessage(message string) {
    r.Message = message
}

func (r *RDSDBSnapshotEvent) SetSourceType(sourceType string) {
    r.SourceType = sourceType
}

func (r *RDSDBSnapshotEvent) SetSourceIdentifier(sourceIdentifier string) {
    r.SourceIdentifier = sourceIdentifier
}

func (r *RDSDBSnapshotEvent) SetEventCategories(eventCategories []string) {
    r.EventCategories = eventCategories
}

func (r *RDSDBSnapshotEvent) SetDate(date time.Time) {
    r.Date = date
}

func (r *RDSDBSnapshotEvent) SetTags(tags map[string]string) {
    r.Tags = tags
}
