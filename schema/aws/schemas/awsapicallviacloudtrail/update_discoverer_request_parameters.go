package awsapicallviacloudtrail

type UpdateDiscovererRequestParameters struct {
    DiscovererId string `json:"discovererId,omitempty"`
    CrossAccount bool `json:"CrossAccount,omitempty"`
    Description string `json:"Description,omitempty"`
}

func (u *UpdateDiscovererRequestParameters) SetDiscovererId(discovererId string) {
    u.DiscovererId = discovererId
}

func (u *UpdateDiscovererRequestParameters) SetCrossAccount(crossAccount bool) {
    u.CrossAccount = crossAccount
}

func (u *UpdateDiscovererRequestParameters) SetDescription(description string) {
    u.Description = description
}
