package awsapicallviacloudtrail

type DescribeDiscovererRequestParameters struct {
    DiscovererId string `json:"discovererId"`
}

func (d *DescribeDiscovererRequestParameters) SetDiscovererId(discovererId string) {
    d.DiscovererId = discovererId
}
