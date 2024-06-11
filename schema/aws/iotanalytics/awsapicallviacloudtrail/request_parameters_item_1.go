package awsapicallviacloudtrail

type RequestParametersItem_1 struct {
    Key string `json:"key"`
    Value string `json:"value"`
}

func (r *RequestParametersItem_1) SetKey(key string) {
    r.Key = key
}

func (r *RequestParametersItem_1) SetValue(value string) {
    r.Value = value
}
