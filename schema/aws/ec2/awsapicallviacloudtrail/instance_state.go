package awsapicallviacloudtrail

type InstanceState struct {
    Code float64 `json:"code,omitempty"`
    Name string `json:"name,omitempty"`
}

func (i *InstanceState) SetCode(code float64) {
    i.Code = code
}

func (i *InstanceState) SetName(name string) {
    i.Name = name
}
