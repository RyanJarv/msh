package sagemakermodelbuildingpipelineexecutionstepstatuschange

import (
    "time"
)


type SageMakerModelBuildingPipelineExecutionStepStatusChange struct {
    CacheHitResult CacheHitResult `json:"cacheHitResult"`
    Metadata Metadata `json:"metadata"`
    CurrentStepStatus string `json:"currentStepStatus"`
    FailureReason string `json:"failureReason"`
    PipelineArn string `json:"pipelineArn"`
    PipelineExecutionArn string `json:"pipelineExecutionArn"`
    PreviousStepStatus string `json:"previousStepStatus"`
    StepEndTime time.Time `json:"stepEndTime"`
    StepName string `json:"stepName"`
    StepStartTime time.Time `json:"stepStartTime"`
    StepType string `json:"stepType"`
}

func (s *SageMakerModelBuildingPipelineExecutionStepStatusChange) SetCacheHitResult(cacheHitResult CacheHitResult) {
    s.CacheHitResult = cacheHitResult
}

func (s *SageMakerModelBuildingPipelineExecutionStepStatusChange) SetMetadata(metadata Metadata) {
    s.Metadata = metadata
}

func (s *SageMakerModelBuildingPipelineExecutionStepStatusChange) SetCurrentStepStatus(currentStepStatus string) {
    s.CurrentStepStatus = currentStepStatus
}

func (s *SageMakerModelBuildingPipelineExecutionStepStatusChange) SetFailureReason(failureReason string) {
    s.FailureReason = failureReason
}

func (s *SageMakerModelBuildingPipelineExecutionStepStatusChange) SetPipelineArn(pipelineArn string) {
    s.PipelineArn = pipelineArn
}

func (s *SageMakerModelBuildingPipelineExecutionStepStatusChange) SetPipelineExecutionArn(pipelineExecutionArn string) {
    s.PipelineExecutionArn = pipelineExecutionArn
}

func (s *SageMakerModelBuildingPipelineExecutionStepStatusChange) SetPreviousStepStatus(previousStepStatus string) {
    s.PreviousStepStatus = previousStepStatus
}

func (s *SageMakerModelBuildingPipelineExecutionStepStatusChange) SetStepEndTime(stepEndTime time.Time) {
    s.StepEndTime = stepEndTime
}

func (s *SageMakerModelBuildingPipelineExecutionStepStatusChange) SetStepName(stepName string) {
    s.StepName = stepName
}

func (s *SageMakerModelBuildingPipelineExecutionStepStatusChange) SetStepStartTime(stepStartTime time.Time) {
    s.StepStartTime = stepStartTime
}

func (s *SageMakerModelBuildingPipelineExecutionStepStatusChange) SetStepType(stepType string) {
    s.StepType = stepType
}
