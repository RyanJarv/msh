package securityhubfindingsimported

type SecurityHubFindingsImported struct {
    Findings interface{} `json:"findings"`
}

func (s *SecurityHubFindingsImported) SetFindings(findings interface{}) {
    s.Findings = findings
}
