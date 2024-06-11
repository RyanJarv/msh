package syntheticscanarytestrunfailure

type Canary_run_timeline struct {
    Completed float64 `json:"completed"`
    Started float64 `json:"started"`
}

func (c *Canary_run_timeline) SetCompleted(completed float64) {
    c.Completed = completed
}

func (c *Canary_run_timeline) SetStarted(started float64) {
    c.Started = started
}
