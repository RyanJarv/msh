package codebuildbuildphasechange

type CodeBuildBuildPhaseChange struct {
    AdditionalInformation Additional_information `json:"additional-information"`
    BuildId string `json:"build-id"`
    CompletedPhase string `json:"completed-phase"`
    CompletedPhaseContext string `json:"completed-phase-context"`
    CompletedPhaseDurationSeconds float64 `json:"completed-phase-duration-seconds"`
    CompletedPhaseEnd string `json:"completed-phase-end"`
    CompletedPhaseStart string `json:"completed-phase-start"`
    CompletedPhaseStatus string `json:"completed-phase-status"`
    ProjectName string `json:"project-name"`
    Version string `json:"version"`
}

func (c *CodeBuildBuildPhaseChange) SetAdditionalInformation(additionalInformation Additional_information) {
    c.AdditionalInformation = additionalInformation
}

func (c *CodeBuildBuildPhaseChange) SetBuildId(buildId string) {
    c.BuildId = buildId
}

func (c *CodeBuildBuildPhaseChange) SetCompletedPhase(completedPhase string) {
    c.CompletedPhase = completedPhase
}

func (c *CodeBuildBuildPhaseChange) SetCompletedPhaseContext(completedPhaseContext string) {
    c.CompletedPhaseContext = completedPhaseContext
}

func (c *CodeBuildBuildPhaseChange) SetCompletedPhaseDurationSeconds(completedPhaseDurationSeconds float64) {
    c.CompletedPhaseDurationSeconds = completedPhaseDurationSeconds
}

func (c *CodeBuildBuildPhaseChange) SetCompletedPhaseEnd(completedPhaseEnd string) {
    c.CompletedPhaseEnd = completedPhaseEnd
}

func (c *CodeBuildBuildPhaseChange) SetCompletedPhaseStart(completedPhaseStart string) {
    c.CompletedPhaseStart = completedPhaseStart
}

func (c *CodeBuildBuildPhaseChange) SetCompletedPhaseStatus(completedPhaseStatus string) {
    c.CompletedPhaseStatus = completedPhaseStatus
}

func (c *CodeBuildBuildPhaseChange) SetProjectName(projectName string) {
    c.ProjectName = projectName
}

func (c *CodeBuildBuildPhaseChange) SetVersion(version string) {
    c.Version = version
}
