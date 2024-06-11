package codebuildbuildstatechange

type Vpc_configItem struct {
    BuildFleetAz string `json:"build-fleet-az"`
    CustomerAz string `json:"customer-az"`
    SubnetId string `json:"subnet-id"`
}

func (v *Vpc_configItem) SetBuildFleetAz(buildFleetAz string) {
    v.BuildFleetAz = buildFleetAz
}

func (v *Vpc_configItem) SetCustomerAz(customerAz string) {
    v.CustomerAz = customerAz
}

func (v *Vpc_configItem) SetSubnetId(subnetId string) {
    v.SubnetId = subnetId
}
