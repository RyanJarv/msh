package awsapicallviacloudtrail

type Math struct {
    Attribute string `json:"attribute,omitempty"`
    Math string `json:"math,omitempty"`
    Name string `json:"name,omitempty"`
}

func (m *Math) SetAttribute(attribute string) {
    m.Attribute = attribute
}

func (m *Math) SetMath(math string) {
    m.Math = math
}

func (m *Math) SetName(name string) {
    m.Name = name
}
