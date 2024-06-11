package ec2fleetstatechange

type EC2FleetStateChange struct {
    SubType string `json:"sub-type"`
}

func (e *EC2FleetStateChange) SetSubType(subType string) {
    e.SubType = subType
}
