package awsapicallviacloudtrail

type AwsvpcConfiguration struct {
    AssignPublicIp string `json:"assignPublicIp,omitempty"`
    Subnets []string `json:"subnets,omitempty"`
}

func (a *AwsvpcConfiguration) SetAssignPublicIp(assignPublicIp string) {
    a.AssignPublicIp = assignPublicIp
}

func (a *AwsvpcConfiguration) SetSubnets(subnets []string) {
    a.Subnets = subnets
}
