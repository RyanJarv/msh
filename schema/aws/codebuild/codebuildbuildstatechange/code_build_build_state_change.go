package codebuildbuildstatechange

type CodeBuildBuildStateChange struct {
    AdditionalInformation Additional_information `json:"additional-information"`
    BuildId string `json:"build-id"`
    BuildStatus string `json:"build-status"`
    CurrentPhase string `json:"current-phase"`
    CurrentPhaseContext string `json:"current-phase-context"`
    ProjectName string `json:"project-name"`
    Version string `json:"version"`
}

func (c *CodeBuildBuildStateChange) SetAdditionalInformation(additionalInformation Additional_information) {
    c.AdditionalInformation = additionalInformation
}

func (c *CodeBuildBuildStateChange) SetBuildId(buildId string) {
    c.BuildId = buildId
}

func (c *CodeBuildBuildStateChange) SetBuildStatus(buildStatus string) {
    c.BuildStatus = buildStatus
}

func (c *CodeBuildBuildStateChange) SetCurrentPhase(currentPhase string) {
    c.CurrentPhase = currentPhase
}

func (c *CodeBuildBuildStateChange) SetCurrentPhaseContext(currentPhaseContext string) {
    c.CurrentPhaseContext = currentPhaseContext
}

func (c *CodeBuildBuildStateChange) SetProjectName(projectName string) {
    c.ProjectName = projectName
}

func (c *CodeBuildBuildStateChange) SetVersion(version string) {
    c.Version = version
}
