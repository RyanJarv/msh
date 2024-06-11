package awsapicallviacloudtrail

type CapacityReservationSpecification struct {
    CapacityReservationPreference string `json:"capacityReservationPreference,omitempty"`
}

func (c *CapacityReservationSpecification) SetCapacityReservationPreference(capacityReservationPreference string) {
    c.CapacityReservationPreference = capacityReservationPreference
}
