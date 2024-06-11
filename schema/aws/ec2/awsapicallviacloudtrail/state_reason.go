package awsapicallviacloudtrail

type StateReason struct {
    Code string `json:"code,omitempty"`
    Message string `json:"message,omitempty"`
}

func (s *StateReason) SetCode(code string) {
    s.Code = code
}

func (s *StateReason) SetMessage(message string) {
    s.Message = message
}
