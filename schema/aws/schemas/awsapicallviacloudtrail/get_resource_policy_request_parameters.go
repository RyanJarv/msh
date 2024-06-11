package awsapicallviacloudtrail

type GetResourcePolicyRequestParameters struct {
    RegistryName string `json:"registryName,omitempty"`
}

func (g *GetResourcePolicyRequestParameters) SetRegistryName(registryName string) {
    g.RegistryName = registryName
}
