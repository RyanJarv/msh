package sagemakermodelbuildingpipelineexecutionstatuschange

import (
    "time"
)


type SageMakerModelBuildingPipelineExecutionStatusChange struct {
    CurrentPipelineExecutionStatus string `json:"currentPipelineExecutionStatus"`
    ExecutionEndTime time.Time `json:"executionEndTime"`
    ExecutionStartTime time.Time `json:"executionStartTime"`
    PipelineArn string `json:"pipelineArn"`
    PipelineExecutionArn string `json:"pipelineExecutionArn"`
    PipelineExecutionDescription string `json:"pipelineExecutionDescription"`
    PipelineExecutionDisplayName string `json:"pipelineExecutionDisplayName"`
    PreviousPipelineExecutionStatus string `json:"previousPipelineExecutionStatus"`
}

func (s *SageMakerModelBuildingPipelineExecutionStatusChange) SetCurrentPipelineExecutionStatus(currentPipelineExecutionStatus string) {
    s.CurrentPipelineExecutionStatus = currentPipelineExecutionStatus
}

func (s *SageMakerModelBuildingPipelineExecutionStatusChange) SetExecutionEndTime(executionEndTime time.Time) {
    s.ExecutionEndTime = executionEndTime
}

func (s *SageMakerModelBuildingPipelineExecutionStatusChange) SetExecutionStartTime(executionStartTime time.Time) {
    s.ExecutionStartTime = executionStartTime
}

func (s *SageMakerModelBuildingPipelineExecutionStatusChange) SetPipelineArn(pipelineArn string) {
    s.PipelineArn = pipelineArn
}

func (s *SageMakerModelBuildingPipelineExecutionStatusChange) SetPipelineExecutionArn(pipelineExecutionArn string) {
    s.PipelineExecutionArn = pipelineExecutionArn
}

func (s *SageMakerModelBuildingPipelineExecutionStatusChange) SetPipelineExecutionDescription(pipelineExecutionDescription string) {
    s.PipelineExecutionDescription = pipelineExecutionDescription
}

func (s *SageMakerModelBuildingPipelineExecutionStatusChange) SetPipelineExecutionDisplayName(pipelineExecutionDisplayName string) {
    s.PipelineExecutionDisplayName = pipelineExecutionDisplayName
}

func (s *SageMakerModelBuildingPipelineExecutionStatusChange) SetPreviousPipelineExecutionStatus(previousPipelineExecutionStatus string) {
    s.PreviousPipelineExecutionStatus = previousPipelineExecutionStatus
}
