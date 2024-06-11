package stepfunctionsexecutionstatuschange

type StepFunctionsExecutionStatusChange struct {
    Output string `json:"output"`
    Input string `json:"input"`
    ExecutionArn string `json:"executionArn"`
    Name string `json:"name"`
    StateMachineArn string `json:"stateMachineArn"`
    StartDate int64 `json:"startDate"`
    StopDate int64 `json:"stopDate"`
    Status string `json:"status"`
}

func (s *StepFunctionsExecutionStatusChange) SetOutput(output string) {
    s.Output = output
}

func (s *StepFunctionsExecutionStatusChange) SetInput(input string) {
    s.Input = input
}

func (s *StepFunctionsExecutionStatusChange) SetExecutionArn(executionArn string) {
    s.ExecutionArn = executionArn
}

func (s *StepFunctionsExecutionStatusChange) SetName(name string) {
    s.Name = name
}

func (s *StepFunctionsExecutionStatusChange) SetStateMachineArn(stateMachineArn string) {
    s.StateMachineArn = stateMachineArn
}

func (s *StepFunctionsExecutionStatusChange) SetStartDate(startDate int64) {
    s.StartDate = startDate
}

func (s *StepFunctionsExecutionStatusChange) SetStopDate(stopDate int64) {
    s.StopDate = stopDate
}

func (s *StepFunctionsExecutionStatusChange) SetStatus(status string) {
    s.Status = status
}
