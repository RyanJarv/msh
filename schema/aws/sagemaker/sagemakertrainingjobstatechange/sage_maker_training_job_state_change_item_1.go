package sagemakertrainingjobstatechange

type SageMakerTrainingJobStateChangeItem_1 struct {
    Status string `json:"Status,omitempty"`
    StartTime float64 `json:"StartTime,omitempty"`
    StatusMessage string `json:"StatusMessage,omitempty"`
}

func (s *SageMakerTrainingJobStateChangeItem_1) SetStatus(status string) {
    s.Status = status
}

func (s *SageMakerTrainingJobStateChangeItem_1) SetStartTime(startTime float64) {
    s.StartTime = startTime
}

func (s *SageMakerTrainingJobStateChangeItem_1) SetStatusMessage(statusMessage string) {
    s.StatusMessage = statusMessage
}
