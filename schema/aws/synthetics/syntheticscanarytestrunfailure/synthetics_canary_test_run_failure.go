package syntheticscanarytestrunfailure

type SyntheticsCanaryTestRunFailure struct {
    CanaryRunTimeline Canary_run_timeline `json:"canary-run-timeline"`
    AccountId string `json:"account-id"`
    ArtifactLocation string `json:"artifact-location"`
    CanaryId string `json:"canary-id"`
    CanaryName string `json:"canary-name"`
    CanaryRunId string `json:"canary-run-id"`
    Message string `json:"message"`
    StateReason string `json:"state-reason"`
    TestRunStatus string `json:"test-run-status"`
}

func (s *SyntheticsCanaryTestRunFailure) SetCanaryRunTimeline(canaryRunTimeline Canary_run_timeline) {
    s.CanaryRunTimeline = canaryRunTimeline
}

func (s *SyntheticsCanaryTestRunFailure) SetAccountId(accountId string) {
    s.AccountId = accountId
}

func (s *SyntheticsCanaryTestRunFailure) SetArtifactLocation(artifactLocation string) {
    s.ArtifactLocation = artifactLocation
}

func (s *SyntheticsCanaryTestRunFailure) SetCanaryId(canaryId string) {
    s.CanaryId = canaryId
}

func (s *SyntheticsCanaryTestRunFailure) SetCanaryName(canaryName string) {
    s.CanaryName = canaryName
}

func (s *SyntheticsCanaryTestRunFailure) SetCanaryRunId(canaryRunId string) {
    s.CanaryRunId = canaryRunId
}

func (s *SyntheticsCanaryTestRunFailure) SetMessage(message string) {
    s.Message = message
}

func (s *SyntheticsCanaryTestRunFailure) SetStateReason(stateReason string) {
    s.StateReason = stateReason
}

func (s *SyntheticsCanaryTestRunFailure) SetTestRunStatus(testRunStatus string) {
    s.TestRunStatus = testRunStatus
}
