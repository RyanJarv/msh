package awsapicallviacloudtrail

type SpotOptions_1 struct {
    SpotInstanceType string `json:"spotInstanceType,omitempty"`
    MaxPrice string `json:"maxPrice,omitempty"`
}

func (s *SpotOptions_1) SetSpotInstanceType(spotInstanceType string) {
    s.SpotInstanceType = spotInstanceType
}

func (s *SpotOptions_1) SetMaxPrice(maxPrice string) {
    s.MaxPrice = maxPrice
}
