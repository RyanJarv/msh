package rdsdbinstanceevent

import (
    "time"
)


type RDSDBInstanceEvent struct {
    Date time.Time `json:"Date,omitempty"`
    EventCategories []string `json:"EventCategories,omitempty"`
    EventID string `json:"EventID,omitempty"`
    Message string `json:"Message,omitempty"`
    SourceArn string `json:"SourceArn,omitempty"`
    SourceIdentifier string `json:"SourceIdentifier,omitempty"`
    SourceType string `json:"SourceType,omitempty"`
    Tags map[string]string `json:"Tags,omitempty"`
}

func (r *RDSDBInstanceEvent) SetDate(date time.Time) {
    r.Date = date
}

func (r *RDSDBInstanceEvent) SetEventCategories(eventCategories []string) {
    r.EventCategories = eventCategories
}

func (r *RDSDBInstanceEvent) SetEventID(eventID string) {
    r.EventID = eventID
}

func (r *RDSDBInstanceEvent) SetMessage(message string) {
    r.Message = message
}

func (r *RDSDBInstanceEvent) SetSourceArn(sourceArn string) {
    r.SourceArn = sourceArn
}

func (r *RDSDBInstanceEvent) SetSourceIdentifier(sourceIdentifier string) {
    r.SourceIdentifier = sourceIdentifier
}

func (r *RDSDBInstanceEvent) SetSourceType(sourceType string) {
    r.SourceType = sourceType
}

func (r *RDSDBInstanceEvent) SetTags(tags map[string]string) {
    r.Tags = tags
}
