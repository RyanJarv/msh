package awsapicallviacloudtrail

type PolicyDetailsItemItem struct {
    Key string `json:"Key"`
    Value string `json:"Value"`
}

func (p *PolicyDetailsItemItem) SetKey(key string) {
    p.Key = key
}

func (p *PolicyDetailsItemItem) SetValue(value string) {
    p.Value = value
}
