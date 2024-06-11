package awsapicallviacloudtrail

type DescribeRegistryRequestParameters struct {
    RegistryName string `json:"registryName"`
}

func (d *DescribeRegistryRequestParameters) SetRegistryName(registryName string) {
    d.RegistryName = registryName
}
