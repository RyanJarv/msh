package ec2instanceterminateunsuccessful

type Details struct {
    SubnetID string `json:"Subnet ID"`
    AvailabilityZone string `json:"Availability Zone"`
}

func (d *Details) SetSubnetID(subnetID string) {
    d.SubnetID = subnetID
}

func (d *Details) SetAvailabilityZone(availabilityZone string) {
    d.AvailabilityZone = availabilityZone
}
