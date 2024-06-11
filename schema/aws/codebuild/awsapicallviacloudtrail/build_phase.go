package awsapicallviacloudtrail

type BuildPhase struct {
    DurationInSeconds float64 `json:"durationInSeconds,omitempty"`
    EndTime string `json:"endTime,omitempty"`
    PhaseContext []string `json:"phaseContext,omitempty"`
    PhaseStatus string `json:"phaseStatus,omitempty"`
    PhaseType string `json:"phaseType,omitempty"`
    StartTime string `json:"startTime,omitempty"`
}

func (b *BuildPhase) SetDurationInSeconds(durationInSeconds float64) {
    b.DurationInSeconds = durationInSeconds
}

func (b *BuildPhase) SetEndTime(endTime string) {
    b.EndTime = endTime
}

func (b *BuildPhase) SetPhaseContext(phaseContext []string) {
    b.PhaseContext = phaseContext
}

func (b *BuildPhase) SetPhaseStatus(phaseStatus string) {
    b.PhaseStatus = phaseStatus
}

func (b *BuildPhase) SetPhaseType(phaseType string) {
    b.PhaseType = phaseType
}

func (b *BuildPhase) SetStartTime(startTime string) {
    b.StartTime = startTime
}
