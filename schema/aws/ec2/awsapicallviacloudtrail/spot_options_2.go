package awsapicallviacloudtrail

type SpotOptions_2 struct {
    SpotInstanceType string `json:"SpotInstanceType,omitempty"`
    MaxPrice float64 `json:"MaxPrice,omitempty"`
}

func (s *SpotOptions_2) SetSpotInstanceType(spotInstanceType string) {
    s.SpotInstanceType = spotInstanceType
}

func (s *SpotOptions_2) SetMaxPrice(maxPrice float64) {
    s.MaxPrice = maxPrice
}
