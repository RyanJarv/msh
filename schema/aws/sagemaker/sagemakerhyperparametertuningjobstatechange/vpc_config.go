package sagemakerhyperparametertuningjobstatechange

type VpcConfig struct {
    Subnets []string `json:"Subnets"`
    SecurityGroupIds []string `json:"SecurityGroupIds"`
}

func (v *VpcConfig) SetSubnets(subnets []string) {
    v.Subnets = subnets
}

func (v *VpcConfig) SetSecurityGroupIds(securityGroupIds []string) {
    v.SecurityGroupIds = securityGroupIds
}
