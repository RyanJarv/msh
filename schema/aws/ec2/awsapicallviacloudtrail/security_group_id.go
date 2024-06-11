package awsapicallviacloudtrail

type SecurityGroupId struct {
    Tag float64 `json:"tag,omitempty"`
    Content string `json:"content,omitempty"`
}

func (s *SecurityGroupId) SetTag(tag float64) {
    s.Tag = tag
}

func (s *SecurityGroupId) SetContent(content string) {
    s.Content = content
}
