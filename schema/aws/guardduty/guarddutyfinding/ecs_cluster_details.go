package guarddutyfinding

type EcsClusterDetails struct {
    TaskDetails TaskDetails `json:"taskDetails,omitempty"`
    Arn string `json:"arn,omitempty"`
    Name string `json:"name,omitempty"`
    Status string `json:"status,omitempty"`
    Tags []EcsClusterDetailsItem `json:"tags,omitempty"`
}

func (e *EcsClusterDetails) SetTaskDetails(taskDetails TaskDetails) {
    e.TaskDetails = taskDetails
}

func (e *EcsClusterDetails) SetArn(arn string) {
    e.Arn = arn
}

func (e *EcsClusterDetails) SetName(name string) {
    e.Name = name
}

func (e *EcsClusterDetails) SetStatus(status string) {
    e.Status = status
}

func (e *EcsClusterDetails) SetTags(tags []EcsClusterDetailsItem) {
    e.Tags = tags
}
