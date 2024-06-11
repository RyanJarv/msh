package securityhubfindingscustomaction

type SecurityHubFindingsCustomAction struct {
    Findings interface{} `json:"findings"`
}

func (s *SecurityHubFindingsCustomAction) SetFindings(findings interface{}) {
    s.Findings = findings
}
