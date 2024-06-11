package awsapicallviacloudtrail

type RequestParametersItem struct {
    Name string `json:"name"`
    Value string `json:"value"`
}

func (r *RequestParametersItem) SetName(name string) {
    r.Name = name
}

func (r *RequestParametersItem) SetValue(value string) {
    r.Value = value
}
