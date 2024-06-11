package awsapicallviacloudtrail

type SessionContext struct {
    WebIdFederationData interface{} `json:"webIdFederationData,omitempty"`
    SessionIssuer SessionIssuer `json:"sessionIssuer,omitempty"`
    Attributes Attributes `json:"attributes,omitempty"`
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
