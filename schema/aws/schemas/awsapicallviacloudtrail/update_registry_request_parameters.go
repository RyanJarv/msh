package awsapicallviacloudtrail

type UpdateRegistryRequestParameters struct {
    RegistryName string `json:"registryName"`
    Description string `json:"Description,omitempty"`
}

func (u *UpdateRegistryRequestParameters) SetRegistryName(registryName string) {
    u.RegistryName = registryName
}

func (u *UpdateRegistryRequestParameters) SetDescription(description string) {
    u.Description = description
}
