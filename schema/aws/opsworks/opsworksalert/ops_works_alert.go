package opsworksalert

type OpsWorksAlert struct {
    StackId string `json:"stack-id"`
    InstanceId string `json:"instance-id"`
    OpsWorksAlertType string `json:"type"`
    Message string `json:"message"`
}

func (o *OpsWorksAlert) SetStackId(stackId string) {
    o.StackId = stackId
}

func (o *OpsWorksAlert) SetInstanceId(instanceId string) {
    o.InstanceId = instanceId
}

func (o *OpsWorksAlert) SetOpsWorksAlertType(opsWorksAlertType string) {
    o.OpsWorksAlertType = opsWorksAlertType
}

func (o *OpsWorksAlert) SetMessage(message string) {
    o.Message = message
}
