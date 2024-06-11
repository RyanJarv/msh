package ecscontainerinstancestatechange

import (
    "time"
)


type ECSContainerInstanceStateChange struct {
    VersionInfo VersionInfo `json:"versionInfo,omitempty"`
    Ec2InstanceId string `json:"ec2InstanceId,omitempty"`
    Attachments []AttachmentDetails `json:"attachments,omitempty"`
    RegisteredResources []ResourceDetails `json:"registeredResources,omitempty"`
    RemainingResources []ResourceDetails `json:"remainingResources,omitempty"`
    RunningTasksCount float64 `json:"runningTasksCount,omitempty"`
    RegisteredAt time.Time `json:"registeredAt,omitempty"`
    AgentConnected bool `json:"agentConnected,omitempty"`
    AgentUpdateStatus string `json:"agentUpdateStatus,omitempty"`
    Version float64 `json:"version,omitempty"`
    PendingTasksCount float64 `json:"pendingTasksCount,omitempty"`
    ClusterArn string `json:"clusterArn,omitempty"`
    Attributes []AttributesDetails `json:"attributes,omitempty"`
    ContainerInstanceArn string `json:"containerInstanceArn,omitempty"`
    Status string `json:"status,omitempty"`
    StatusReason string `json:"statusReason,omitempty"`
    UpdatedAt time.Time `json:"updatedAt,omitempty"`
    AccountType string `json:"accountType,omitempty"`
}

func (e *ECSContainerInstanceStateChange) SetVersionInfo(versionInfo VersionInfo) {
    e.VersionInfo = versionInfo
}

func (e *ECSContainerInstanceStateChange) SetEc2InstanceId(ec2InstanceId string) {
    e.Ec2InstanceId = ec2InstanceId
}

func (e *ECSContainerInstanceStateChange) SetAttachments(attachments []AttachmentDetails) {
    e.Attachments = attachments
}

func (e *ECSContainerInstanceStateChange) SetRegisteredResources(registeredResources []ResourceDetails) {
    e.RegisteredResources = registeredResources
}

func (e *ECSContainerInstanceStateChange) SetRemainingResources(remainingResources []ResourceDetails) {
    e.RemainingResources = remainingResources
}

func (e *ECSContainerInstanceStateChange) SetRunningTasksCount(runningTasksCount float64) {
    e.RunningTasksCount = runningTasksCount
}

func (e *ECSContainerInstanceStateChange) SetRegisteredAt(registeredAt time.Time) {
    e.RegisteredAt = registeredAt
}

func (e *ECSContainerInstanceStateChange) SetAgentConnected(agentConnected bool) {
    e.AgentConnected = agentConnected
}

func (e *ECSContainerInstanceStateChange) SetAgentUpdateStatus(agentUpdateStatus string) {
    e.AgentUpdateStatus = agentUpdateStatus
}

func (e *ECSContainerInstanceStateChange) SetVersion(version float64) {
    e.Version = version
}

func (e *ECSContainerInstanceStateChange) SetPendingTasksCount(pendingTasksCount float64) {
    e.PendingTasksCount = pendingTasksCount
}

func (e *ECSContainerInstanceStateChange) SetClusterArn(clusterArn string) {
    e.ClusterArn = clusterArn
}

func (e *ECSContainerInstanceStateChange) SetAttributes(attributes []AttributesDetails) {
    e.Attributes = attributes
}

func (e *ECSContainerInstanceStateChange) SetContainerInstanceArn(containerInstanceArn string) {
    e.ContainerInstanceArn = containerInstanceArn
}

func (e *ECSContainerInstanceStateChange) SetStatus(status string) {
    e.Status = status
}

func (e *ECSContainerInstanceStateChange) SetStatusReason(statusReason string) {
    e.StatusReason = statusReason
}

func (e *ECSContainerInstanceStateChange) SetUpdatedAt(updatedAt time.Time) {
    e.UpdatedAt = updatedAt
}

func (e *ECSContainerInstanceStateChange) SetAccountType(accountType string) {
    e.AccountType = accountType
}
