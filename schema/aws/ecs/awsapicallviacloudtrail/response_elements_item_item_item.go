package awsapicallviacloudtrail

type ResponseElementsItemItemItem struct {
    Name string `json:"name"`
    Value string `json:"value"`
}

func (r *ResponseElementsItemItemItem) SetName(name string) {
    r.Name = name
}

func (r *ResponseElementsItemItemItem) SetValue(value string) {
    r.Value = value
}
