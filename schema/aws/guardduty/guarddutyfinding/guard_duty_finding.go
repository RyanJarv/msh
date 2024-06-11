package guarddutyfinding

import (
    "time"
)


type GuardDutyFinding struct {
    Resource Resource `json:"resource"`
    Service Service `json:"service"`
    AccountId string `json:"accountId"`
    Arn string `json:"arn"`
    CreatedAt time.Time `json:"createdAt"`
    Description string `json:"description"`
    Id string `json:"id"`
    Partition string `json:"partition"`
    Region string `json:"region"`
    SchemaVersion string `json:"schemaVersion"`
    Severity float64 `json:"severity"`
    Title string `json:"title"`
    GuardDutyFindingType string `json:"type"`
    UpdatedAt time.Time `json:"updatedAt"`
}

func (g *GuardDutyFinding) SetResource(resource Resource) {
    g.Resource = resource
}

func (g *GuardDutyFinding) SetService(service Service) {
    g.Service = service
}

func (g *GuardDutyFinding) SetAccountId(accountId string) {
    g.AccountId = accountId
}

func (g *GuardDutyFinding) SetArn(arn string) {
    g.Arn = arn
}

func (g *GuardDutyFinding) SetCreatedAt(createdAt time.Time) {
    g.CreatedAt = createdAt
}

func (g *GuardDutyFinding) SetDescription(description string) {
    g.Description = description
}

func (g *GuardDutyFinding) SetId(id string) {
    g.Id = id
}

func (g *GuardDutyFinding) SetPartition(partition string) {
    g.Partition = partition
}

func (g *GuardDutyFinding) SetRegion(region string) {
    g.Region = region
}

func (g *GuardDutyFinding) SetSchemaVersion(schemaVersion string) {
    g.SchemaVersion = schemaVersion
}

func (g *GuardDutyFinding) SetSeverity(severity float64) {
    g.Severity = severity
}

func (g *GuardDutyFinding) SetTitle(title string) {
    g.Title = title
}

func (g *GuardDutyFinding) SetGuardDutyFindingType(guardDutyFindingType string) {
    g.GuardDutyFindingType = guardDutyFindingType
}

func (g *GuardDutyFinding) SetUpdatedAt(updatedAt time.Time) {
    g.UpdatedAt = updatedAt
}
