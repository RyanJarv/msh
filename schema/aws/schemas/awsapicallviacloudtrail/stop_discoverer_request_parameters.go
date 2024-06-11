package awsapicallviacloudtrail

type StopDiscovererRequestParameters struct {
    DiscovererId string `json:"discovererId"`
}

func (s *StopDiscovererRequestParameters) SetDiscovererId(discovererId string) {
    s.DiscovererId = discovererId
}
