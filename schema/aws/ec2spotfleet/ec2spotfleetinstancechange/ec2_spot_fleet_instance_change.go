package ec2spotfleetinstancechange

type EC2SpotFleetInstanceChange struct {
    Description string `json:"description"`
    InstanceId string `json:"instance-id"`
    SubType string `json:"sub-type"`
}

func (e *EC2SpotFleetInstanceChange) SetDescription(description string) {
    e.Description = description
}

func (e *EC2SpotFleetInstanceChange) SetInstanceId(instanceId string) {
    e.InstanceId = instanceId
}

func (e *EC2SpotFleetInstanceChange) SetSubType(subType string) {
    e.SubType = subType
}
