package sagemakerhyperparametertuningjobstatechange

import (
    "time"
)


type SageMakerHyperParameterTuningJobStateChange struct {
    TrainingJobStatusCounters TrainingJobStatusCounters `json:"TrainingJobStatusCounters"`
    ObjectiveStatusCounters ObjectiveStatusCounters `json:"ObjectiveStatusCounters"`
    TrainingJobDefinition TrainingJobDefinition `json:"TrainingJobDefinition"`
    Tags []Tags `json:"Tags,omitempty"`
    CreationTime time.Time `json:"CreationTime"`
    LastModifiedTime time.Time `json:"LastModifiedTime"`
    HyperParameterTuningJobName string `json:"HyperParameterTuningJobName"`
    HyperParameterTuningJobStatus string `json:"HyperParameterTuningJobStatus"`
    HyperParameterTuningJobArn string `json:"HyperParameterTuningJobArn"`
}

func (s *SageMakerHyperParameterTuningJobStateChange) SetTrainingJobStatusCounters(trainingJobStatusCounters TrainingJobStatusCounters) {
    s.TrainingJobStatusCounters = trainingJobStatusCounters
}

func (s *SageMakerHyperParameterTuningJobStateChange) SetObjectiveStatusCounters(objectiveStatusCounters ObjectiveStatusCounters) {
    s.ObjectiveStatusCounters = objectiveStatusCounters
}

func (s *SageMakerHyperParameterTuningJobStateChange) SetTrainingJobDefinition(trainingJobDefinition TrainingJobDefinition) {
    s.TrainingJobDefinition = trainingJobDefinition
}

func (s *SageMakerHyperParameterTuningJobStateChange) SetTags(tags []Tags) {
    s.Tags = tags
}

func (s *SageMakerHyperParameterTuningJobStateChange) SetCreationTime(creationTime time.Time) {
    s.CreationTime = creationTime
}

func (s *SageMakerHyperParameterTuningJobStateChange) SetLastModifiedTime(lastModifiedTime time.Time) {
    s.LastModifiedTime = lastModifiedTime
}

func (s *SageMakerHyperParameterTuningJobStateChange) SetHyperParameterTuningJobName(hyperParameterTuningJobName string) {
    s.HyperParameterTuningJobName = hyperParameterTuningJobName
}

func (s *SageMakerHyperParameterTuningJobStateChange) SetHyperParameterTuningJobStatus(hyperParameterTuningJobStatus string) {
    s.HyperParameterTuningJobStatus = hyperParameterTuningJobStatus
}

func (s *SageMakerHyperParameterTuningJobStateChange) SetHyperParameterTuningJobArn(hyperParameterTuningJobArn string) {
    s.HyperParameterTuningJobArn = hyperParameterTuningJobArn
}
