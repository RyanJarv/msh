package awsapicallviacloudtrail

type ProjectItem struct {
    Key string `json:"key,omitempty"`
    Value string `json:"value,omitempty"`
}

func (p *ProjectItem) SetKey(key string) {
    p.Key = key
}

func (p *ProjectItem) SetValue(value string) {
    p.Value = value
}
