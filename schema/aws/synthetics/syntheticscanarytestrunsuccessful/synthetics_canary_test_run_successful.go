package syntheticscanarytestrunsuccessful

type SyntheticsCanaryTestRunSuccessful struct {
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

func (s *SyntheticsCanaryTestRunSuccessful) SetCanaryRunTimeline(canaryRunTimeline Canary_run_timeline) {
    s.CanaryRunTimeline = canaryRunTimeline
}

func (s *SyntheticsCanaryTestRunSuccessful) SetAccountId(accountId string) {
    s.AccountId = accountId
}

func (s *SyntheticsCanaryTestRunSuccessful) SetArtifactLocation(artifactLocation string) {
    s.ArtifactLocation = artifactLocation
}

func (s *SyntheticsCanaryTestRunSuccessful) SetCanaryId(canaryId string) {
    s.CanaryId = canaryId
}

func (s *SyntheticsCanaryTestRunSuccessful) SetCanaryName(canaryName string) {
    s.CanaryName = canaryName
}

func (s *SyntheticsCanaryTestRunSuccessful) SetCanaryRunId(canaryRunId string) {
    s.CanaryRunId = canaryRunId
}

func (s *SyntheticsCanaryTestRunSuccessful) SetMessage(message string) {
    s.Message = message
}

func (s *SyntheticsCanaryTestRunSuccessful) SetStateReason(stateReason string) {
    s.StateReason = stateReason
}

func (s *SyntheticsCanaryTestRunSuccessful) SetTestRunStatus(testRunStatus string) {
    s.TestRunStatus = testRunStatus
}
