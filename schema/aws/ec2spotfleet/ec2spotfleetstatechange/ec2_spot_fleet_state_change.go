package ec2spotfleetstatechange

type EC2SpotFleetStateChange struct {
    SubType string `json:"sub-type"`
}

func (e *EC2SpotFleetStateChange) SetSubType(subType string) {
    e.SubType = subType
}
