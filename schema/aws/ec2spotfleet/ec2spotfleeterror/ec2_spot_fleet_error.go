package ec2spotfleeterror

type EC2SpotFleetError struct {
    Description string `json:"description"`
    SubType string `json:"sub-type"`
}

func (e *EC2SpotFleetError) SetDescription(description string) {
    e.Description = description
}

func (e *EC2SpotFleetError) SetSubType(subType string) {
    e.SubType = subType
}
