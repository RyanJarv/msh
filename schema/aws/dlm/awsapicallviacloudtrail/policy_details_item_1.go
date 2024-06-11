package awsapicallviacloudtrail

type PolicyDetailsItem_1 struct {
    Key string `json:"Key"`
    Value string `json:"Value"`
}

func (p *PolicyDetailsItem_1) SetKey(key string) {
    p.Key = key
}

func (p *PolicyDetailsItem_1) SetValue(value string) {
    p.Value = value
}
