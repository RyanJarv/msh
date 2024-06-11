package ec2spotfleetinformation

type EC2SpotFleetInformation struct {
    Description string `json:"description"`
    SubType string `json:"sub-type"`
}

func (e *EC2SpotFleetInformation) SetDescription(description string) {
    e.Description = description
}

func (e *EC2SpotFleetInformation) SetSubType(subType string) {
    e.SubType = subType
}
