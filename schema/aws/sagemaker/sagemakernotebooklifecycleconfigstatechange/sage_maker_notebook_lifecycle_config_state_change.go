package sagemakernotebooklifecycleconfigstatechange

type SageMakerNotebookLifecycleConfigStateChange struct {
    CreationTime float64 `json:"CreationTime,omitempty"`
    LastModifiedTime float64 `json:"LastModifiedTime,omitempty"`
    NotebookInstanceLifecycleConfigArn string `json:"NotebookInstanceLifecycleConfigArn,omitempty"`
    NotebookInstanceLifecycleConfigName string `json:"NotebookInstanceLifecycleConfigName,omitempty"`
    Tags []Tags `json:"Tags,omitempty"`
}

func (s *SageMakerNotebookLifecycleConfigStateChange) SetCreationTime(creationTime float64) {
    s.CreationTime = creationTime
}

func (s *SageMakerNotebookLifecycleConfigStateChange) SetLastModifiedTime(lastModifiedTime float64) {
    s.LastModifiedTime = lastModifiedTime
}

func (s *SageMakerNotebookLifecycleConfigStateChange) SetNotebookInstanceLifecycleConfigArn(notebookInstanceLifecycleConfigArn string) {
    s.NotebookInstanceLifecycleConfigArn = notebookInstanceLifecycleConfigArn
}

func (s *SageMakerNotebookLifecycleConfigStateChange) SetNotebookInstanceLifecycleConfigName(notebookInstanceLifecycleConfigName string) {
    s.NotebookInstanceLifecycleConfigName = notebookInstanceLifecycleConfigName
}

func (s *SageMakerNotebookLifecycleConfigStateChange) SetTags(tags []Tags) {
    s.Tags = tags
}
