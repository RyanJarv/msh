package codebuildbuildphasechange

type Additional_informationItem struct {
    DurationInSeconds float64 `json:"duration-in-seconds,omitempty"`
    EndTime string `json:"end-time,omitempty"`
    PhaseContext []string `json:"phase-context,omitempty"`
    PhaseStatus string `json:"phase-status,omitempty"`
    PhaseType string `json:"phase-type"`
    StartTime string `json:"start-time"`
}

func (a *Additional_informationItem) SetDurationInSeconds(durationInSeconds float64) {
    a.DurationInSeconds = durationInSeconds
}

func (a *Additional_informationItem) SetEndTime(endTime string) {
    a.EndTime = endTime
}

func (a *Additional_informationItem) SetPhaseContext(phaseContext []string) {
    a.PhaseContext = phaseContext
}

func (a *Additional_informationItem) SetPhaseStatus(phaseStatus string) {
    a.PhaseStatus = phaseStatus
}

func (a *Additional_informationItem) SetPhaseType(phaseType string) {
    a.PhaseType = phaseType
}

func (a *Additional_informationItem) SetStartTime(startTime string) {
    a.StartTime = startTime
}
