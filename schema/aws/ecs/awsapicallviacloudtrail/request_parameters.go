package awsapicallviacloudtrail

type RequestParameters struct {
    NetworkConfiguration NetworkConfiguration `json:"networkConfiguration,omitempty"`
    Overrides Overrides `json:"overrides,omitempty"`
    PlacementConstraints []RequestParametersItem `json:"placementConstraints,omitempty"`
    Cluster string `json:"cluster"`
    Reason string `json:"reason,omitempty"`
    ExecutionStoppedAt string `json:"executionStoppedAt,omitempty"`
    PullStoppedAt string `json:"pullStoppedAt,omitempty"`
    StartedBy string `json:"startedBy,omitempty"`
    PullStartedAt string `json:"pullStartedAt,omitempty"`
    Count float64 `json:"count,omitempty"`
    Task string `json:"task,omitempty"`
    ContainerInstance string `json:"containerInstance,omitempty"`
    Containers []RequestParametersItem_1 `json:"containers,omitempty"`
    TaskDefinition string `json:"taskDefinition,omitempty"`
    EnableECSManagedTags bool `json:"enableECSManagedTags,omitempty"`
    LaunchType string `json:"launchType,omitempty"`
    Status string `json:"status,omitempty"`
}

func (r *RequestParameters) SetNetworkConfiguration(networkConfiguration NetworkConfiguration) {
    r.NetworkConfiguration = networkConfiguration
}

func (r *RequestParameters) SetOverrides(overrides Overrides) {
    r.Overrides = overrides
}

func (r *RequestParameters) SetPlacementConstraints(placementConstraints []RequestParametersItem) {
    r.PlacementConstraints = placementConstraints
}

func (r *RequestParameters) SetCluster(cluster string) {
    r.Cluster = cluster
}

func (r *RequestParameters) SetReason(reason string) {
    r.Reason = reason
}

func (r *RequestParameters) SetExecutionStoppedAt(executionStoppedAt string) {
    r.ExecutionStoppedAt = executionStoppedAt
}

func (r *RequestParameters) SetPullStoppedAt(pullStoppedAt string) {
    r.PullStoppedAt = pullStoppedAt
}

func (r *RequestParameters) SetStartedBy(startedBy string) {
    r.StartedBy = startedBy
}

func (r *RequestParameters) SetPullStartedAt(pullStartedAt string) {
    r.PullStartedAt = pullStartedAt
}

func (r *RequestParameters) SetCount(count float64) {
    r.Count = count
}

func (r *RequestParameters) SetTask(task string) {
    r.Task = task
}

func (r *RequestParameters) SetContainerInstance(containerInstance string) {
    r.ContainerInstance = containerInstance
}

func (r *RequestParameters) SetContainers(containers []RequestParametersItem_1) {
    r.Containers = containers
}

func (r *RequestParameters) SetTaskDefinition(taskDefinition string) {
    r.TaskDefinition = taskDefinition
}

func (r *RequestParameters) SetEnableECSManagedTags(enableECSManagedTags bool) {
    r.EnableECSManagedTags = enableECSManagedTags
}

func (r *RequestParameters) SetLaunchType(launchType string) {
    r.LaunchType = launchType
}

func (r *RequestParameters) SetStatus(status string) {
    r.Status = status
}
