package awsapicallviacloudtrail

type SessionContext struct {
    WebIdFederationData interface{} `json:"webIdFederationData"`
    SessionIssuer SessionIssuer `json:"sessionIssuer"`
    Attributes Attributes `json:"attributes"`
}

func (s *SessionContext) SetWebIdFederationData(webIdFederationData interface{}) {
    s.WebIdFederationData = webIdFederationData
}

func (s *SessionContext) SetSessionIssuer(sessionIssuer SessionIssuer) {
    s.SessionIssuer = sessionIssuer
}

func (s *SessionContext) SetAttributes(attributes Attributes) {
    s.Attributes = attributes
}
