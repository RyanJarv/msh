package awsapicallviacloudtrail

type NetworkConfiguration struct {
    AwsvpcConfiguration AwsvpcConfiguration `json:"awsvpcConfiguration,omitempty"`
}

func (n *NetworkConfiguration) SetAwsvpcConfiguration(awsvpcConfiguration AwsvpcConfiguration) {
    n.AwsvpcConfiguration = awsvpcConfiguration
}
