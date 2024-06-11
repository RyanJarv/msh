package simpleworkflowexecutionstatechange

type WorkflowExecutionDetail struct {
    ExecutionConfiguration ExecutionConfiguration `json:"executionConfiguration"`
    ExecutionInfo ExecutionInfo `json:"executionInfo"`
    OpenCounts OpenCounts `json:"openCounts"`
    LatestActivityTaskTimestamp interface{} `json:"latestActivityTaskTimestamp"`
}

func (w *WorkflowExecutionDetail) SetExecutionConfiguration(executionConfiguration ExecutionConfiguration) {
    w.ExecutionConfiguration = executionConfiguration
}

func (w *WorkflowExecutionDetail) SetExecutionInfo(executionInfo ExecutionInfo) {
    w.ExecutionInfo = executionInfo
}

func (w *WorkflowExecutionDetail) SetOpenCounts(openCounts OpenCounts) {
    w.OpenCounts = openCounts
}

func (w *WorkflowExecutionDetail) SetLatestActivityTaskTimestamp(latestActivityTaskTimestamp interface{}) {
    w.LatestActivityTaskTimestamp = latestActivityTaskTimestamp
}
