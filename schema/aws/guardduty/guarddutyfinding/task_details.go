package guarddutyfinding

type TaskDetails struct {
    Arn string `json:"arn,omitempty"`
    Containers []TaskDetailsItem `json:"containers,omitempty"`
    CreatedAt float64 `json:"createdAt,omitempty"`
    DefinitionArn string `json:"definitionArn,omitempty"`
    StartedAt float64 `json:"startedAt,omitempty"`
    StartedBy string `json:"startedBy,omitempty"`
    Version string `json:"version,omitempty"`
}

func (t *TaskDetails) SetArn(arn string) {
    t.Arn = arn
}

func (t *TaskDetails) SetContainers(containers []TaskDetailsItem) {
    t.Containers = containers
}

func (t *TaskDetails) SetCreatedAt(createdAt float64) {
    t.CreatedAt = createdAt
}

func (t *TaskDetails) SetDefinitionArn(definitionArn string) {
    t.DefinitionArn = definitionArn
}

func (t *TaskDetails) SetStartedAt(startedAt float64) {
    t.StartedAt = startedAt
}

func (t *TaskDetails) SetStartedBy(startedBy string) {
    t.StartedBy = startedBy
}

func (t *TaskDetails) SetVersion(version string) {
    t.Version = version
}
