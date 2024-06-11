package awsapicallviacloudtrail

type HealthCheck struct {
    UnhealthyThreshold float64 `json:"unhealthyThreshold,omitempty"`
    Interval float64 `json:"interval,omitempty"`
    HealthyThreshold float64 `json:"healthyThreshold,omitempty"`
    Timeout float64 `json:"timeout,omitempty"`
    Target string `json:"target,omitempty"`
}

func (h *HealthCheck) SetUnhealthyThreshold(unhealthyThreshold float64) {
    h.UnhealthyThreshold = unhealthyThreshold
}

func (h *HealthCheck) SetInterval(interval float64) {
    h.Interval = interval
}

func (h *HealthCheck) SetHealthyThreshold(healthyThreshold float64) {
    h.HealthyThreshold = healthyThreshold
}

func (h *HealthCheck) SetTimeout(timeout float64) {
    h.Timeout = timeout
}

func (h *HealthCheck) SetTarget(target string) {
    h.Target = target
}
