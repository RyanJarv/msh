package awsapicallviacloudtrail

type InstanceMarketOptions_1 struct {
    SpotOptions SpotOptions_2 `json:"SpotOptions,omitempty"`
    MarketType string `json:"MarketType,omitempty"`
}

func (i *InstanceMarketOptions_1) SetSpotOptions(spotOptions SpotOptions_2) {
    i.SpotOptions = spotOptions
}

func (i *InstanceMarketOptions_1) SetMarketType(marketType string) {
    i.MarketType = marketType
}
