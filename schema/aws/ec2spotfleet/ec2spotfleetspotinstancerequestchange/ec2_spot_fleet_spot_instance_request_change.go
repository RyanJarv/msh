package ec2spotfleetspotinstancerequestchange

type EC2SpotFleetSpotInstanceRequestChange struct {
    Description string `json:"description"`
    SpotInstanceRequestId string `json:"spot-instance-request-id"`
    SubType string `json:"sub-type"`
}

func (e *EC2SpotFleetSpotInstanceRequestChange) SetDescription(description string) {
    e.Description = description
}

func (e *EC2SpotFleetSpotInstanceRequestChange) SetSpotInstanceRequestId(spotInstanceRequestId string) {
    e.SpotInstanceRequestId = spotInstanceRequestId
}

func (e *EC2SpotFleetSpotInstanceRequestChange) SetSubType(subType string) {
    e.SubType = subType
}
