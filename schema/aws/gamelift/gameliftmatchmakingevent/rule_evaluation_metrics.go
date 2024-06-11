package gameliftmatchmakingevent

type RuleEvaluationMetrics struct {
    FailedCount float64 `json:"failedCount"`
    PassedCount float64 `json:"passedCount"`
    RuleName string `json:"ruleName"`
}

func (r *RuleEvaluationMetrics) SetFailedCount(failedCount float64) {
    r.FailedCount = failedCount
}

func (r *RuleEvaluationMetrics) SetPassedCount(passedCount float64) {
    r.PassedCount = passedCount
}

func (r *RuleEvaluationMetrics) SetRuleName(ruleName string) {
    r.RuleName = ruleName
}
