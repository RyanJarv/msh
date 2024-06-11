package awsapicallviacloudtrail

type SessionContext struct {
    Attributes Attributes `json:"attributes"`
    SessionIssuer SessionIssuer `json:"sessionIssuer"`
    WebIdFederationData WebIdFederationData `json:"webIdFederationData"`
}

func (s *SessionContext) SetAttributes(attributes Attributes) {
    s.Attributes = attributes
}

func (s *SessionContext) SetSessionIssuer(sessionIssuer SessionIssuer) {
    s.SessionIssuer = sessionIssuer
}

func (s *SessionContext) SetWebIdFederationData(webIdFederationData WebIdFederationData) {
    s.WebIdFederationData = webIdFederationData
}
