package awsapicallviacloudtrail

type HealthCheck_1 struct {
    UnhealthyThreshold float64 `json:"unhealthyThreshold,omitempty"`
    Interval float64 `json:"interval,omitempty"`
    HealthyThreshold float64 `json:"healthyThreshold,omitempty"`
    Timeout float64 `json:"timeout,omitempty"`
    Target string `json:"target,omitempty"`
}

func (h *HealthCheck_1) SetUnhealthyThreshold(unhealthyThreshold float64) {
    h.UnhealthyThreshold = unhealthyThreshold
}

func (h *HealthCheck_1) SetInterval(interval float64) {
    h.Interval = interval
}

func (h *HealthCheck_1) SetHealthyThreshold(healthyThreshold float64) {
    h.HealthyThreshold = healthyThreshold
}

func (h *HealthCheck_1) SetTimeout(timeout float64) {
    h.Timeout = timeout
}

func (h *HealthCheck_1) SetTarget(target string) {
    h.Target = target
}
