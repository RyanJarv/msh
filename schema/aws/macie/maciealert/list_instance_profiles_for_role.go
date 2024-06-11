package maciealert

import (
    "time"
)


type ListInstanceProfilesForRole struct {
    Multiplier float64 `json:"multiplier,omitempty"`
    Narrative string `json:"narrative,omitempty"`
    Name string `json:"name,omitempty"`
    Description string `json:"description,omitempty"`
    Risk float64 `json:"risk,omitempty"`
    Anomalous bool `json:"anomalous,omitempty"`
    ExcessionTimes []time.Time `json:"excession_times,omitempty"`
}

func (l *ListInstanceProfilesForRole) SetMultiplier(multiplier float64) {
    l.Multiplier = multiplier
}

func (l *ListInstanceProfilesForRole) SetNarrative(narrative string) {
    l.Narrative = narrative
}

func (l *ListInstanceProfilesForRole) SetName(name string) {
    l.Name = name
}

func (l *ListInstanceProfilesForRole) SetDescription(description string) {
    l.Description = description
}

func (l *ListInstanceProfilesForRole) SetRisk(risk float64) {
    l.Risk = risk
}

func (l *ListInstanceProfilesForRole) SetAnomalous(anomalous bool) {
    l.Anomalous = anomalous
}

func (l *ListInstanceProfilesForRole) SetExcessionTimes(excessionTimes []time.Time) {
    l.ExcessionTimes = excessionTimes
}
