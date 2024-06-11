package ec2instancelaunchunsuccessful

type Details struct {
    SubnetID string `json:"Subnet ID,omitempty"`
    AvailabilityZone string `json:"Availability Zone,omitempty"`
}

func (d *Details) SetSubnetID(subnetID string) {
    d.SubnetID = subnetID
}

func (d *Details) SetAvailabilityZone(availabilityZone string) {
    d.AvailabilityZone = availabilityZone
}
