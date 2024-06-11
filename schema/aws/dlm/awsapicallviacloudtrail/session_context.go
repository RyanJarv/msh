package awsapicallviacloudtrail

type SessionContext struct {
    Attributes Attributes `json:"attributes"`
    SessionIssuer SessionIssuer `json:"sessionIssuer"`
}

func (s *SessionContext) SetAttributes(attributes Attributes) {
    s.Attributes = attributes
}

func (s *SessionContext) SetSessionIssuer(sessionIssuer SessionIssuer) {
    s.SessionIssuer = sessionIssuer
}
