package simpleworkflowexecutionstatechange

type ExecutionConfiguration struct {
    TaskList TaskList `json:"taskList"`
    ChildPolicy string `json:"childPolicy"`
    ExecutionStartToCloseTimeout string `json:"executionStartToCloseTimeout"`
    LambdaRole string `json:"lambdaRole"`
    TaskPriority interface{} `json:"taskPriority"`
    TaskStartToCloseTimeout string `json:"taskStartToCloseTimeout"`
}

func (e *ExecutionConfiguration) SetTaskList(taskList TaskList) {
    e.TaskList = taskList
}

func (e *ExecutionConfiguration) SetChildPolicy(childPolicy string) {
    e.ChildPolicy = childPolicy
}

func (e *ExecutionConfiguration) SetExecutionStartToCloseTimeout(executionStartToCloseTimeout string) {
    e.ExecutionStartToCloseTimeout = executionStartToCloseTimeout
}

func (e *ExecutionConfiguration) SetLambdaRole(lambdaRole string) {
    e.LambdaRole = lambdaRole
}

func (e *ExecutionConfiguration) SetTaskPriority(taskPriority interface{}) {
    e.TaskPriority = taskPriority
}

func (e *ExecutionConfiguration) SetTaskStartToCloseTimeout(taskStartToCloseTimeout string) {
    e.TaskStartToCloseTimeout = taskStartToCloseTimeout
}
