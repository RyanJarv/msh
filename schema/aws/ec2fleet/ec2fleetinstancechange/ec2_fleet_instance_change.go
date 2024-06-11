package ec2fleetinstancechange

type EC2FleetInstanceChange struct {
    Description string `json:"description"`
    InstanceId string `json:"instance-id"`
    SubType string `json:"sub-type"`
}

func (e *EC2FleetInstanceChange) SetDescription(description string) {
    e.Description = description
}

func (e *EC2FleetInstanceChange) SetInstanceId(instanceId string) {
    e.InstanceId = instanceId
}

func (e *EC2FleetInstanceChange) SetSubType(subType string) {
    e.SubType = subType
}
