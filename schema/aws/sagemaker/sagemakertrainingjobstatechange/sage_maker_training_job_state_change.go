package sagemakertrainingjobstatechange

type SageMakerTrainingJobStateChange struct {
    HyperParameters HyperParameters `json:"HyperParameters,omitempty"`
    AlgorithmSpecification AlgorithmSpecification `json:"AlgorithmSpecification,omitempty"`
    OutputDataConfig OutputDataConfig `json:"OutputDataConfig,omitempty"`
    ModelArtifacts ModelArtifacts `json:"ModelArtifacts,omitempty"`
    StoppingCondition StoppingCondition `json:"StoppingCondition,omitempty"`
    ResourceConfig ResourceConfig `json:"ResourceConfig,omitempty"`
    Tags []Tags `json:"Tags,omitempty"`
    TrainingJobName string `json:"TrainingJobName,omitempty"`
    SecondaryStatus string `json:"SecondaryStatus,omitempty"`
    InputDataConfig []SageMakerTrainingJobStateChangeItem `json:"InputDataConfig,omitempty"`
    SecondaryStatusTransitions []SageMakerTrainingJobStateChangeItem_1 `json:"SecondaryStatusTransitions,omitempty"`
    RoleArn string `json:"RoleArn,omitempty"`
    TrainingJobStatus string `json:"TrainingJobStatus,omitempty"`
    TrainingJobArn string `json:"TrainingJobArn,omitempty"`
    CreationTime float64 `json:"CreationTime,omitempty"`
    LastModifiedTime float64 `json:"LastModifiedTime,omitempty"`
    TrainingStartTime float64 `json:"TrainingStartTime,omitempty"`
    TrainingEndTime float64 `json:"TrainingEndTime,omitempty"`
}

func (s *SageMakerTrainingJobStateChange) SetHyperParameters(hyperParameters HyperParameters) {
    s.HyperParameters = hyperParameters
}

func (s *SageMakerTrainingJobStateChange) SetAlgorithmSpecification(algorithmSpecification AlgorithmSpecification) {
    s.AlgorithmSpecification = algorithmSpecification
}

func (s *SageMakerTrainingJobStateChange) SetOutputDataConfig(outputDataConfig OutputDataConfig) {
    s.OutputDataConfig = outputDataConfig
}

func (s *SageMakerTrainingJobStateChange) SetModelArtifacts(modelArtifacts ModelArtifacts) {
    s.ModelArtifacts = modelArtifacts
}

func (s *SageMakerTrainingJobStateChange) SetStoppingCondition(stoppingCondition StoppingCondition) {
    s.StoppingCondition = stoppingCondition
}

func (s *SageMakerTrainingJobStateChange) SetResourceConfig(resourceConfig ResourceConfig) {
    s.ResourceConfig = resourceConfig
}

func (s *SageMakerTrainingJobStateChange) SetTags(tags []Tags) {
    s.Tags = tags
}

func (s *SageMakerTrainingJobStateChange) SetTrainingJobName(trainingJobName string) {
    s.TrainingJobName = trainingJobName
}

func (s *SageMakerTrainingJobStateChange) SetSecondaryStatus(secondaryStatus string) {
    s.SecondaryStatus = secondaryStatus
}

func (s *SageMakerTrainingJobStateChange) SetInputDataConfig(inputDataConfig []SageMakerTrainingJobStateChangeItem) {
    s.InputDataConfig = inputDataConfig
}

func (s *SageMakerTrainingJobStateChange) SetSecondaryStatusTransitions(secondaryStatusTransitions []SageMakerTrainingJobStateChangeItem_1) {
    s.SecondaryStatusTransitions = secondaryStatusTransitions
}

func (s *SageMakerTrainingJobStateChange) SetRoleArn(roleArn string) {
    s.RoleArn = roleArn
}

func (s *SageMakerTrainingJobStateChange) SetTrainingJobStatus(trainingJobStatus string) {
    s.TrainingJobStatus = trainingJobStatus
}

func (s *SageMakerTrainingJobStateChange) SetTrainingJobArn(trainingJobArn string) {
    s.TrainingJobArn = trainingJobArn
}

func (s *SageMakerTrainingJobStateChange) SetCreationTime(creationTime float64) {
    s.CreationTime = creationTime
}

func (s *SageMakerTrainingJobStateChange) SetLastModifiedTime(lastModifiedTime float64) {
    s.LastModifiedTime = lastModifiedTime
}

func (s *SageMakerTrainingJobStateChange) SetTrainingStartTime(trainingStartTime float64) {
    s.TrainingStartTime = trainingStartTime
}

func (s *SageMakerTrainingJobStateChange) SetTrainingEndTime(trainingEndTime float64) {
    s.TrainingEndTime = trainingEndTime
}
