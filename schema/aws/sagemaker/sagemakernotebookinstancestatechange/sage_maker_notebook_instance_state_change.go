package sagemakernotebookinstancestatechange

type SageMakerNotebookInstanceStateChange struct {
    Tags []Tags `json:"Tags,omitempty"`
    CreationTime float64 `json:"CreationTime,omitempty"`
    InstanceType string `json:"InstanceType,omitempty"`
    LastModifiedTime float64 `json:"LastModifiedTime,omitempty"`
    NotebookInstanceArn string `json:"NotebookInstanceArn,omitempty"`
    NotebookInstanceLifecycleConfigName string `json:"NotebookInstanceLifecycleConfigName,omitempty"`
    NotebookInstanceName string `json:"NotebookInstanceName,omitempty"`
    NotebookInstanceStatus string `json:"NotebookInstanceStatus,omitempty"`
    RoleArn string `json:"RoleArn,omitempty"`
}

func (s *SageMakerNotebookInstanceStateChange) SetTags(tags []Tags) {
    s.Tags = tags
}

func (s *SageMakerNotebookInstanceStateChange) SetCreationTime(creationTime float64) {
    s.CreationTime = creationTime
}

func (s *SageMakerNotebookInstanceStateChange) SetInstanceType(instanceType string) {
    s.InstanceType = instanceType
}

func (s *SageMakerNotebookInstanceStateChange) SetLastModifiedTime(lastModifiedTime float64) {
    s.LastModifiedTime = lastModifiedTime
}

func (s *SageMakerNotebookInstanceStateChange) SetNotebookInstanceArn(notebookInstanceArn string) {
    s.NotebookInstanceArn = notebookInstanceArn
}

func (s *SageMakerNotebookInstanceStateChange) SetNotebookInstanceLifecycleConfigName(notebookInstanceLifecycleConfigName string) {
    s.NotebookInstanceLifecycleConfigName = notebookInstanceLifecycleConfigName
}

func (s *SageMakerNotebookInstanceStateChange) SetNotebookInstanceName(notebookInstanceName string) {
    s.NotebookInstanceName = notebookInstanceName
}

func (s *SageMakerNotebookInstanceStateChange) SetNotebookInstanceStatus(notebookInstanceStatus string) {
    s.NotebookInstanceStatus = notebookInstanceStatus
}

func (s *SageMakerNotebookInstanceStateChange) SetRoleArn(roleArn string) {
    s.RoleArn = roleArn
}
