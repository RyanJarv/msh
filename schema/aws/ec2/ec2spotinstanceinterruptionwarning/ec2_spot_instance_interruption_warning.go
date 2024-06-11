package ec2spotinstanceinterruptionwarning

type EC2SpotInstanceInterruptionWarning struct {
    InstanceId string `json:"instance-id"`
    InstanceAction string `json:"instance-action"`
}

func (e *EC2SpotInstanceInterruptionWarning) SetInstanceId(instanceId string) {
    e.InstanceId = instanceId
}

func (e *EC2SpotInstanceInterruptionWarning) SetInstanceAction(instanceAction string) {
    e.InstanceAction = instanceAction
}
