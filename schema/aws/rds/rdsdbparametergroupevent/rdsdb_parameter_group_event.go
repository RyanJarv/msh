package rdsdbparametergroupevent

import (
    "time"
)


type RDSDBParameterGroupEvent struct {
    SourceArn string `json:"SourceArn,omitempty"`
    Message string `json:"Message,omitempty"`
    SourceType string `json:"SourceType,omitempty"`
    SourceIdentifier string `json:"SourceIdentifier,omitempty"`
    EventCategories []string `json:"EventCategories,omitempty"`
    Date time.Time `json:"Date,omitempty"`
    Tags map[string]string `json:"Tags,omitempty"`
}

func (r *RDSDBParameterGroupEvent) SetSourceArn(sourceArn string) {
    r.SourceArn = sourceArn
}

func (r *RDSDBParameterGroupEvent) SetMessage(message string) {
    r.Message = message
}

func (r *RDSDBParameterGroupEvent) SetSourceType(sourceType string) {
    r.SourceType = sourceType
}

func (r *RDSDBParameterGroupEvent) SetSourceIdentifier(sourceIdentifier string) {
    r.SourceIdentifier = sourceIdentifier
}

func (r *RDSDBParameterGroupEvent) SetEventCategories(eventCategories []string) {
    r.EventCategories = eventCategories
}

func (r *RDSDBParameterGroupEvent) SetDate(date time.Time) {
    r.Date = date
}

func (r *RDSDBParameterGroupEvent) SetTags(tags map[string]string) {
    r.Tags = tags
}
