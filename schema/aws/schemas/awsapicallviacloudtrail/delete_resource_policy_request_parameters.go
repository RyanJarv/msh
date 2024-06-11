package awsapicallviacloudtrail

type DeleteResourcePolicyRequestParameters struct {
    RegistryName string `json:"registryName,omitempty"`
}

func (d *DeleteResourcePolicyRequestParameters) SetRegistryName(registryName string) {
    d.RegistryName = registryName
}
