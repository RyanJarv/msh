package codeguruprofilerrecommendationstatechange

import (
    "time"
)


type CodeGuruProfilerRecommendationsStateChange struct {
    Recommendation Recommendation `json:"recommendation"`
    Title Title `json:"title"`
    ComputeInstanceArns []string `json:"computeInstanceArns"`
    DeduplicationId string `json:"deduplicationId"`
    EventEndTime time.Time `json:"eventEndTime,omitempty"`
    EventStartTime time.Time `json:"eventStartTime"`
    ExpiresOn time.Time `json:"expiresOn"`
    Schema string `json:"schema"`
    Severity string `json:"severity"`
    SourceUrl string `json:"sourceUrl"`
    Status string `json:"status"`
}

func (c *CodeGuruProfilerRecommendationsStateChange) SetRecommendation(recommendation Recommendation) {
    c.Recommendation = recommendation
}

func (c *CodeGuruProfilerRecommendationsStateChange) SetTitle(title Title) {
    c.Title = title
}

func (c *CodeGuruProfilerRecommendationsStateChange) SetComputeInstanceArns(computeInstanceArns []string) {
    c.ComputeInstanceArns = computeInstanceArns
}

func (c *CodeGuruProfilerRecommendationsStateChange) SetDeduplicationId(deduplicationId string) {
    c.DeduplicationId = deduplicationId
}

func (c *CodeGuruProfilerRecommendationsStateChange) SetEventEndTime(eventEndTime time.Time) {
    c.EventEndTime = eventEndTime
}

func (c *CodeGuruProfilerRecommendationsStateChange) SetEventStartTime(eventStartTime time.Time) {
    c.EventStartTime = eventStartTime
}

func (c *CodeGuruProfilerRecommendationsStateChange) SetExpiresOn(expiresOn time.Time) {
    c.ExpiresOn = expiresOn
}

func (c *CodeGuruProfilerRecommendationsStateChange) SetSchema(schema string) {
    c.Schema = schema
}

func (c *CodeGuruProfilerRecommendationsStateChange) SetSeverity(severity string) {
    c.Severity = severity
}

func (c *CodeGuruProfilerRecommendationsStateChange) SetSourceUrl(sourceUrl string) {
    c.SourceUrl = sourceUrl
}

func (c *CodeGuruProfilerRecommendationsStateChange) SetStatus(status string) {
    c.Status = status
}
