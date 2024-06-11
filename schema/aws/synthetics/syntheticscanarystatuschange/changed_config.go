package syntheticscanarystatuschange

type Changed_config struct {
    ExecutionArn ExecutionArn `json:"executionArn"`
    TestCodeLayerVersionArn TestCodeLayerVersionArn `json:"testCodeLayerVersionArn"`
}

func (c *Changed_config) SetExecutionArn(executionArn ExecutionArn) {
    c.ExecutionArn = executionArn
}

func (c *Changed_config) SetTestCodeLayerVersionArn(testCodeLayerVersionArn TestCodeLayerVersionArn) {
    c.TestCodeLayerVersionArn = testCodeLayerVersionArn
}
