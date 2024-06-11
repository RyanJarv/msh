package awsapicallviacloudtrail

type StartDiscovererRequestParameters struct {
    DiscovererId string `json:"discovererId"`
}

func (s *StartDiscovererRequestParameters) SetDiscovererId(discovererId string) {
    s.DiscovererId = discovererId
}
