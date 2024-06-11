package rdsdbsecuritygroupevent

import (
    "time"
)


type RDSDBSecurityGroupEvent struct {
    SourceArn string `json:"SourceArn,omitempty"`
    Message string `json:"Message,omitempty"`
    SourceType string `json:"SourceType,omitempty"`
    SourceIdentifier string `json:"SourceIdentifier,omitempty"`
    EventCategories []string `json:"EventCategories,omitempty"`
    Date time.Time `json:"Date,omitempty"`
    Tags map[string]string `json:"Tags,omitempty"`
}

func (r *RDSDBSecurityGroupEvent) SetSourceArn(sourceArn string) {
    r.SourceArn = sourceArn
}

func (r *RDSDBSecurityGroupEvent) SetMessage(message string) {
    r.Message = message
}

func (r *RDSDBSecurityGroupEvent) SetSourceType(sourceType string) {
    r.SourceType = sourceType
}

func (r *RDSDBSecurityGroupEvent) SetSourceIdentifier(sourceIdentifier string) {
    r.SourceIdentifier = sourceIdentifier
}

func (r *RDSDBSecurityGroupEvent) SetEventCategories(eventCategories []string) {
    r.EventCategories = eventCategories
}

func (r *RDSDBSecurityGroupEvent) SetDate(date time.Time) {
    r.Date = date
}

func (r *RDSDBSecurityGroupEvent) SetTags(tags map[string]string) {
    r.Tags = tags
}
