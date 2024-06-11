package awsapicallviacloudtrail

type DeleteDiscovererRequestParameters struct {
    DiscovererId string `json:"discovererId"`
}

func (d *DeleteDiscovererRequestParameters) SetDiscovererId(discovererId string) {
    d.DiscovererId = discovererId
}
