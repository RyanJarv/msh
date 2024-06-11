package opsworkscommandstatechange

type OpsWorksCommandStateChange struct {
    CommandId string `json:"command-id"`
    InstanceId string `json:"instance-id"`
    OpsWorksCommandStateChangeType string `json:"type"`
    Status string `json:"status"`
}

func (o *OpsWorksCommandStateChange) SetCommandId(commandId string) {
    o.CommandId = commandId
}

func (o *OpsWorksCommandStateChange) SetInstanceId(instanceId string) {
    o.InstanceId = instanceId
}

func (o *OpsWorksCommandStateChange) SetOpsWorksCommandStateChangeType(opsWorksCommandStateChangeType string) {
    o.OpsWorksCommandStateChangeType = opsWorksCommandStateChangeType
}

func (o *OpsWorksCommandStateChange) SetStatus(status string) {
    o.Status = status
}
