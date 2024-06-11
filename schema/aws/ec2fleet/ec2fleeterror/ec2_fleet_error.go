package ec2fleeterror

type EC2FleetError struct {
    Description string `json:"description"`
    SubType string `json:"sub-type"`
}

func (e *EC2FleetError) SetDescription(description string) {
    e.Description = description
}

func (e *EC2FleetError) SetSubType(subType string) {
    e.SubType = subType
}
