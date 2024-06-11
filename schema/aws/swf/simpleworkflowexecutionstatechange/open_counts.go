package simpleworkflowexecutionstatechange

type OpenCounts struct {
    OpenActivityTasks float64 `json:"openActivityTasks"`
    OpenChildWorkflowExecutions float64 `json:"openChildWorkflowExecutions"`
    OpenDecisionTasks float64 `json:"openDecisionTasks"`
    OpenLambdaFunctions float64 `json:"openLambdaFunctions"`
    OpenTimers float64 `json:"openTimers"`
}

func (o *OpenCounts) SetOpenActivityTasks(openActivityTasks float64) {
    o.OpenActivityTasks = openActivityTasks
}

func (o *OpenCounts) SetOpenChildWorkflowExecutions(openChildWorkflowExecutions float64) {
    o.OpenChildWorkflowExecutions = openChildWorkflowExecutions
}

func (o *OpenCounts) SetOpenDecisionTasks(openDecisionTasks float64) {
    o.OpenDecisionTasks = openDecisionTasks
}

func (o *OpenCounts) SetOpenLambdaFunctions(openLambdaFunctions float64) {
    o.OpenLambdaFunctions = openLambdaFunctions
}

func (o *OpenCounts) SetOpenTimers(openTimers float64) {
    o.OpenTimers = openTimers
}
