package ec2fleetinformation

type EC2FleetInformation struct {
    Description string `json:"description"`
    SubType string `json:"sub-type"`
}

func (e *EC2FleetInformation) SetDescription(description string) {
    e.Description = description
}

func (e *EC2FleetInformation) SetSubType(subType string) {
    e.SubType = subType
}
