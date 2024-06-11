package awsapicallviacloudtrail

type SessionIssuer struct {
    AccountId string `json:"accountId,omitempty"`
    Arn string `json:"arn,omitempty"`
    PrincipalId string `json:"principalId,omitempty"`
    SessionIssuerType string `json:"type,omitempty"`
    UserName string `json:"userName,omitempty"`
}

func (s *SessionIssuer) SetAccountId(accountId string) {
    s.AccountId = accountId
}

func (s *SessionIssuer) SetArn(arn string) {
    s.Arn = arn
}

func (s *SessionIssuer) SetPrincipalId(principalId string) {
    s.PrincipalId = principalId
}

func (s *SessionIssuer) SetSessionIssuerType(sessionIssuerType string) {
    s.SessionIssuerType = sessionIssuerType
}

func (s *SessionIssuer) SetUserName(userName string) {
    s.UserName = userName
}
