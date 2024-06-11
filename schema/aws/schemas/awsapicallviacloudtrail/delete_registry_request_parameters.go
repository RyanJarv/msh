package awsapicallviacloudtrail

type DeleteRegistryRequestParameters struct {
    RegistryName string `json:"registryName"`
}

func (d *DeleteRegistryRequestParameters) SetRegistryName(registryName string) {
    d.RegistryName = registryName
}
