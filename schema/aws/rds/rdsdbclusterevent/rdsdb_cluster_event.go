package rdsdbclusterevent

import (
    "time"
)


type RDSDBClusterEvent struct {
    SourceArn string `json:"SourceArn,omitempty"`
    Message string `json:"Message,omitempty"`
    SourceType string `json:"SourceType,omitempty"`
    SourceIdentifier string `json:"SourceIdentifier,omitempty"`
    EventCategories []string `json:"EventCategories,omitempty"`
    Date time.Time `json:"Date,omitempty"`
    Tags map[string]string `json:"Tags,omitempty"`
}

func (r *RDSDBClusterEvent) SetSourceArn(sourceArn string) {
    r.SourceArn = sourceArn
}

func (r *RDSDBClusterEvent) SetMessage(message string) {
    r.Message = message
}

func (r *RDSDBClusterEvent) SetSourceType(sourceType string) {
    r.SourceType = sourceType
}

func (r *RDSDBClusterEvent) SetSourceIdentifier(sourceIdentifier string) {
    r.SourceIdentifier = sourceIdentifier
}

func (r *RDSDBClusterEvent) SetEventCategories(eventCategories []string) {
    r.EventCategories = eventCategories
}

func (r *RDSDBClusterEvent) SetDate(date time.Time) {
    r.Date = date
}

func (r *RDSDBClusterEvent) SetTags(tags map[string]string) {
    r.Tags = tags
}
