package sagemakerhyperparametertuningjobstatechange

type ObjectiveStatusCounters struct {
    Succeeded float64 `json:"Succeeded"`
    Failed float64 `json:"Failed"`
    Pending float64 `json:"Pending"`
}

func (o *ObjectiveStatusCounters) SetSucceeded(succeeded float64) {
    o.Succeeded = succeeded
}

func (o *ObjectiveStatusCounters) SetFailed(failed float64) {
    o.Failed = failed
}

func (o *ObjectiveStatusCounters) SetPending(pending float64) {
    o.Pending = pending
}
