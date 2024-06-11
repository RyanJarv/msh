package limitbreach

type IVSLimitBreach struct {
    ExceededBy float64 `json:"exceeded_by"`
    Limit string `json:"limit"`
    LimitUnit string `json:"limit_unit"`
    LimitValue float64 `json:"limit_value"`
}

func (i *IVSLimitBreach) SetExceededBy(exceededBy float64) {
    i.ExceededBy = exceededBy
}

func (i *IVSLimitBreach) SetLimit(limit string) {
    i.Limit = limit
}

func (i *IVSLimitBreach) SetLimitUnit(limitUnit string) {
    i.LimitUnit = limitUnit
}

func (i *IVSLimitBreach) SetLimitValue(limitValue float64) {
    i.LimitValue = limitValue
}
