package awsapicallviacloudtrail

type Placement struct {
    Tenancy string `json:"tenancy,omitempty"`
    AvailabilityZone string `json:"availabilityZone,omitempty"`
}

func (p *Placement) SetTenancy(tenancy string) {
    p.Tenancy = tenancy
}

func (p *Placement) SetAvailabilityZone(availabilityZone string) {
    p.AvailabilityZone = availabilityZone
}
