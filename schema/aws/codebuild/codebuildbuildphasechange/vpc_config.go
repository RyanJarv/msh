package codebuildbuildphasechange

type Vpc_config struct {
    SecurityGroupIds []string `json:"security-group-ids"`
    Subnets []Vpc_configItem `json:"subnets"`
    VpcId string `json:"vpc-id"`
}

func (v *Vpc_config) SetSecurityGroupIds(securityGroupIds []string) {
    v.SecurityGroupIds = securityGroupIds
}

func (v *Vpc_config) SetSubnets(subnets []Vpc_configItem) {
    v.Subnets = subnets
}

func (v *Vpc_config) SetVpcId(vpcId string) {
    v.VpcId = vpcId
}
