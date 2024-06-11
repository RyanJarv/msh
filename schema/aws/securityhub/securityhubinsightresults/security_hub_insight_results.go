package securityhubinsightresults

type SecurityHubInsightResults struct {
    InsightName string `json:"insightName"`
    InsightResults []SecurityHubInsightResult `json:"insightResults"`
    ActionDescription string `json:"actionDescription"`
    InsightArn string `json:"insightArn"`
    ResultType string `json:"resultType"`
    ActionName string `json:"actionName"`
}

func (s *SecurityHubInsightResults) SetInsightName(insightName string) {
    s.InsightName = insightName
}

func (s *SecurityHubInsightResults) SetInsightResults(insightResults []SecurityHubInsightResult) {
    s.InsightResults = insightResults
}

func (s *SecurityHubInsightResults) SetActionDescription(actionDescription string) {
    s.ActionDescription = actionDescription
}

func (s *SecurityHubInsightResults) SetInsightArn(insightArn string) {
    s.InsightArn = insightArn
}

func (s *SecurityHubInsightResults) SetResultType(resultType string) {
    s.ResultType = resultType
}

func (s *SecurityHubInsightResults) SetActionName(actionName string) {
    s.ActionName = actionName
}
