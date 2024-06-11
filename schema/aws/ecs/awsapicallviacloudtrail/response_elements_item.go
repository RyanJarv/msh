package awsapicallviacloudtrail

type ResponseElementsItem struct {
    Overrides Overrides_1 `json:"overrides"`
    Memory string `json:"memory"`
    Attachments []ResponseElementsItemItem `json:"attachments"`
    StartedBy string `json:"startedBy,omitempty"`
    TaskArn string `json:"taskArn"`
    Cpu string `json:"cpu"`
    Version float64 `json:"version"`
    Tags []interface{} `json:"tags"`
    CreatedAt string `json:"createdAt"`
    ClusterArn string `json:"clusterArn"`
    TaskDefinitionArn string `json:"taskDefinitionArn"`
    PlatformVersion string `json:"platformVersion,omitempty"`
    Containers []ResponseElementsItemItem_1 `json:"containers"`
    ContainerInstanceArn string `json:"containerInstanceArn,omitempty"`
    LastStatus string `json:"lastStatus"`
    DesiredStatus string `json:"desiredStatus"`
    Group string `json:"group"`
    LaunchType string `json:"launchType"`
}

func (r *ResponseElementsItem) SetOverrides(overrides Overrides_1) {
    r.Overrides = overrides
}

func (r *ResponseElementsItem) SetMemory(memory string) {
    r.Memory = memory
}

func (r *ResponseElementsItem) SetAttachments(attachments []ResponseElementsItemItem) {
    r.Attachments = attachments
}

func (r *ResponseElementsItem) SetStartedBy(startedBy string) {
    r.StartedBy = startedBy
}

func (r *ResponseElementsItem) SetTaskArn(taskArn string) {
    r.TaskArn = taskArn
}

func (r *ResponseElementsItem) SetCpu(cpu string) {
    r.Cpu = cpu
}

func (r *ResponseElementsItem) SetVersion(version float64) {
    r.Version = version
}

func (r *ResponseElementsItem) SetTags(tags []interface{}) {
    r.Tags = tags
}

func (r *ResponseElementsItem) SetCreatedAt(createdAt string) {
    r.CreatedAt = createdAt
}

func (r *ResponseElementsItem) SetClusterArn(clusterArn string) {
    r.ClusterArn = clusterArn
}

func (r *ResponseElementsItem) SetTaskDefinitionArn(taskDefinitionArn string) {
    r.TaskDefinitionArn = taskDefinitionArn
}

func (r *ResponseElementsItem) SetPlatformVersion(platformVersion string) {
    r.PlatformVersion = platformVersion
}

func (r *ResponseElementsItem) SetContainers(containers []ResponseElementsItemItem_1) {
    r.Containers = containers
}

func (r *ResponseElementsItem) SetContainerInstanceArn(containerInstanceArn string) {
    r.ContainerInstanceArn = containerInstanceArn
}

func (r *ResponseElementsItem) SetLastStatus(lastStatus string) {
    r.LastStatus = lastStatus
}

func (r *ResponseElementsItem) SetDesiredStatus(desiredStatus string) {
    r.DesiredStatus = desiredStatus
}

func (r *ResponseElementsItem) SetGroup(group string) {
    r.Group = group
}

func (r *ResponseElementsItem) SetLaunchType(launchType string) {
    r.LaunchType = launchType
}
