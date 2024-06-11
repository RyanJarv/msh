package batchjobstatechange

type Container_1 struct {
    NetworkInterfaces []interface{} `json:"networkInterfaces"`
    LogStreamName string `json:"logStreamName"`
    TaskArn string `json:"taskArn"`
    ExitCode float64 `json:"exitCode"`
    ContainerInstanceArn string `json:"containerInstanceArn"`
}

func (c *Container_1) SetNetworkInterfaces(networkInterfaces []interface{}) {
    c.NetworkInterfaces = networkInterfaces
}

func (c *Container_1) SetLogStreamName(logStreamName string) {
    c.LogStreamName = logStreamName
}

func (c *Container_1) SetTaskArn(taskArn string) {
    c.TaskArn = taskArn
}

func (c *Container_1) SetExitCode(exitCode float64) {
    c.ExitCode = exitCode
}

func (c *Container_1) SetContainerInstanceArn(containerInstanceArn string) {
    c.ContainerInstanceArn = containerInstanceArn
}
