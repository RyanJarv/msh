package awsapicallviacloudtrail

type RequestParametersItem_1 struct {
    ContainerName string `json:"containerName"`
    ExitCode float64 `json:"exitCode,omitempty"`
    NetworkBindings []interface{} `json:"networkBindings"`
    Status string `json:"status"`
}

func (r *RequestParametersItem_1) SetContainerName(containerName string) {
    r.ContainerName = containerName
}

func (r *RequestParametersItem_1) SetExitCode(exitCode float64) {
    r.ExitCode = exitCode
}

func (r *RequestParametersItem_1) SetNetworkBindings(networkBindings []interface{}) {
    r.NetworkBindings = networkBindings
}

func (r *RequestParametersItem_1) SetStatus(status string) {
    r.Status = status
}
