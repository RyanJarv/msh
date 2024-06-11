package batchjobstatechange

type ULimit struct {
    HardLimit float64 `json:"hardLimit,omitempty"`
    SoftLimit float64 `json:"softLimit,omitempty"`
    Name string `json:"name,omitempty"`
}

func (u *ULimit) SetHardLimit(hardLimit float64) {
    u.HardLimit = hardLimit
}

func (u *ULimit) SetSoftLimit(softLimit float64) {
    u.SoftLimit = softLimit
}

func (u *ULimit) SetName(name string) {
    u.Name = name
}
