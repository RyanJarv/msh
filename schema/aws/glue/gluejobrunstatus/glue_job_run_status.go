package gluejobrunstatus

import (
    "time"
)


type GlueJobRunStatus struct {
    NotificationCondition NotificationCondition `json:"notificationCondition"`
    JobName string `json:"jobName"`
    Severity string `json:"severity"`
    State string `json:"state"`
    JobRunId string `json:"jobRunId"`
    Message string `json:"message"`
    StartedOn time.Time `json:"startedOn"`
}

func (g *GlueJobRunStatus) SetNotificationCondition(notificationCondition NotificationCondition) {
    g.NotificationCondition = notificationCondition
}

func (g *GlueJobRunStatus) SetJobName(jobName string) {
    g.JobName = jobName
}

func (g *GlueJobRunStatus) SetSeverity(severity string) {
    g.Severity = severity
}

func (g *GlueJobRunStatus) SetState(state string) {
    g.State = state
}

func (g *GlueJobRunStatus) SetJobRunId(jobRunId string) {
    g.JobRunId = jobRunId
}

func (g *GlueJobRunStatus) SetMessage(message string) {
    g.Message = message
}

func (g *GlueJobRunStatus) SetStartedOn(startedOn time.Time) {
    g.StartedOn = startedOn
}
