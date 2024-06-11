package sagemakermodelstatechange

type SageMakerModelStateChange struct {
    PrimaryContainer PrimaryContainer `json:"PrimaryContainer"`
    Tags []Tags `json:"Tags,omitempty"`
    ModelArn string `json:"ModelArn"`
    ExecutionRoleArn string `json:"ExecutionRoleArn"`
    ModelName string `json:"ModelName"`
    CreationTime float64 `json:"CreationTime"`
}

func (s *SageMakerModelStateChange) SetPrimaryContainer(primaryContainer PrimaryContainer) {
    s.PrimaryContainer = primaryContainer
}

func (s *SageMakerModelStateChange) SetTags(tags []Tags) {
    s.Tags = tags
}

func (s *SageMakerModelStateChange) SetModelArn(modelArn string) {
    s.ModelArn = modelArn
}

func (s *SageMakerModelStateChange) SetExecutionRoleArn(executionRoleArn string) {
    s.ExecutionRoleArn = executionRoleArn
}

func (s *SageMakerModelStateChange) SetModelName(modelName string) {
    s.ModelName = modelName
}

func (s *SageMakerModelStateChange) SetCreationTime(creationTime float64) {
    s.CreationTime = creationTime
}
