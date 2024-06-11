package ec2fleetspotinstancerequestchange

type EC2FleetSpotInstanceRequestChange struct {
    Description string `json:"description"`
    SpotInstanceRequestId string `json:"spot-instance-request-id"`
    SubType string `json:"sub-type"`
}

func (e *EC2FleetSpotInstanceRequestChange) SetDescription(description string) {
    e.Description = description
}

func (e *EC2FleetSpotInstanceRequestChange) SetSpotInstanceRequestId(spotInstanceRequestId string) {
    e.SpotInstanceRequestId = spotInstanceRequestId
}

func (e *EC2FleetSpotInstanceRequestChange) SetSubType(subType string) {
    e.SubType = subType
}
