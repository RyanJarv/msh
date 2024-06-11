package awsapicallviacloudtrail

type VpcConfig struct {
    SecurityGroupIds []string `json:"securityGroupIds,omitempty"`
    Subnets []string `json:"subnets,omitempty"`
    VpcId string `json:"vpcId,omitempty"`
}

func (v *VpcConfig) SetSecurityGroupIds(securityGroupIds []string) {
    v.SecurityGroupIds = securityGroupIds
}

func (v *VpcConfig) SetSubnets(subnets []string) {
    v.Subnets = subnets
}

func (v *VpcConfig) SetVpcId(vpcId string) {
    v.VpcId = vpcId
}
