package awsapicallviacloudtrail

type InstanceMarketOptions struct {
    SpotOptions SpotOptions_1 `json:"spotOptions,omitempty"`
    MarketType string `json:"marketType,omitempty"`
}

func (i *InstanceMarketOptions) SetSpotOptions(spotOptions SpotOptions_1) {
    i.SpotOptions = spotOptions
}

func (i *InstanceMarketOptions) SetMarketType(marketType string) {
    i.MarketType = marketType
}
