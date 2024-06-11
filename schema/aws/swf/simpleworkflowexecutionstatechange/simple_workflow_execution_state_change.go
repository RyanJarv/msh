package simpleworkflowexecutionstatechange

type SimpleWorkflowExecutionStateChange struct {
    WorkflowExecutionDetail WorkflowExecutionDetail `json:"workflowExecutionDetail"`
    EventId float64 `json:"eventId"`
    EventType string `json:"eventType"`
}

func (s *SimpleWorkflowExecutionStateChange) SetWorkflowExecutionDetail(workflowExecutionDetail WorkflowExecutionDetail) {
    s.WorkflowExecutionDetail = workflowExecutionDetail
}

func (s *SimpleWorkflowExecutionStateChange) SetEventId(eventId float64) {
    s.EventId = eventId
}

func (s *SimpleWorkflowExecutionStateChange) SetEventType(eventType string) {
    s.EventType = eventType
}
